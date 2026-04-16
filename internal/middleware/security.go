package middleware

import (
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

const maxTrackedVisitors = 100_000

// clientIP extracts the real client IP from proxy headers, falling back to RemoteAddr.
// Honors Fly-Client-IP (Fly.io) and X-Forwarded-For (standard reverse proxies).
func clientIP(r *http.Request) string {
	if ip := strings.TrimSpace(r.Header.Get("Fly-Client-IP")); ip != "" {
		return ip
	}
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		if comma := strings.Index(xff, ","); comma != -1 {
			return strings.TrimSpace(xff[:comma])
		}
		return strings.TrimSpace(xff)
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

// SecurityHeadersMiddleware adds security headers to all responses
func SecurityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set security headers
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

		next.ServeHTTP(w, r)
	})
}

// RateLimiter implements simple token bucket rate limiting per IP
type RateLimiter struct {
	visitors map[string]*visitor
	mu       sync.RWMutex
	rate     int           // requests per second
	burst    int           // burst size
	interval time.Duration // cleanup interval
}

type visitor struct {
	tokens     int
	lastAccess time.Time
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(rate, burst int) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*visitor),
		rate:     rate,
		burst:    burst,
		interval: time.Minute,
	}

	// Start cleanup goroutine
	go rl.cleanup()

	return rl
}

// Middleware returns the rate limiting middleware
func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rl.allow(clientIP(r)) {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// allow checks if the request should be allowed
func (rl *RateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[ip]
	now := time.Now()

	if !exists {
		// Bound memory: if at cap, evict any visitor idle > 1 minute.
		if len(rl.visitors) >= maxTrackedVisitors {
			for k, old := range rl.visitors {
				if now.Sub(old.lastAccess) > time.Minute {
					delete(rl.visitors, k)
				}
			}
			if len(rl.visitors) >= maxTrackedVisitors {
				return false
			}
		}
		rl.visitors[ip] = &visitor{
			tokens:     rl.burst - 1,
			lastAccess: now,
		}
		return true
	}

	// Refill tokens based on time elapsed
	elapsed := now.Sub(v.lastAccess)
	tokensToAdd := int(elapsed.Seconds()) * rl.rate
	v.tokens += tokensToAdd

	if v.tokens > rl.burst {
		v.tokens = rl.burst
	}

	v.lastAccess = now

	if v.tokens > 0 {
		v.tokens--
		return true
	}

	return false
}

// cleanup removes old visitors periodically
func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(rl.interval)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for ip, v := range rl.visitors {
			if now.Sub(v.lastAccess) > 5*time.Minute {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}
