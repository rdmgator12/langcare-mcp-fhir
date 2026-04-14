# KDIGO CKD Staging Reference

## GFR Categories

| Stage | eGFR (mL/min/1.73m2) | Description | Color Risk |
|-------|---------------------|-------------|------------|
| G1 | >=90 | Normal or high | Green (if no albuminuria) |
| G2 | 60-89 | Mildly decreased | Green/Yellow |
| G3a | 45-59 | Mild to moderate decrease | Yellow/Orange |
| G3b | 30-44 | Moderate to severe decrease | Orange/Red |
| G4 | 15-29 | Severely decreased | Red/Dark Red |
| G5 | <15 | Kidney failure | Dark Red |

## Albuminuria Categories

| Category | UACR (mg/g) | AER (mg/24h) | Description |
|----------|-------------|-------------|-------------|
| A1 | <30 | <30 | Normal to mildly increased |
| A2 | 30-300 | 30-300 | Moderately increased |
| A3 | >300 | >300 | Severely increased |

## KDIGO Risk Matrix (GFR x Albuminuria)

```
              A1 (<30)    A2 (30-300)   A3 (>300)
G1 (>=90)     Low Risk    Moderate      High
G2 (60-89)    Low Risk    Moderate      High
G3a (45-59)   Moderate    High          Very High
G3b (30-44)   High        Very High     Very High
G4 (15-29)    Very High   Very High     Very High
G5 (<15)      Very High   Very High     Very High
```

## Monitoring Frequency by Risk

| Risk Level | GFR/UACR Monitoring | Clinical Action |
|-----------|--------------------|-----------------| 
| Low (Green) | Annually | Standard care |
| Moderate (Yellow) | Every 6-12 months | Address risk factors |
| High (Orange) | Every 3-6 months | Nephrology referral if progressing |
| Very High (Red) | Every 1-3 months | Nephrology co-management |

## eGFR Calculation (CKD-EPI 2021 -- Race-Free)

CKD-EPI 2021 equation (no longer uses race coefficient):

**Female (Cr <= 0.7):** 142 x (Cr/0.7)^(-0.241) x 0.9938^age
**Female (Cr > 0.7):** 142 x (Cr/0.7)^(-1.200) x 0.9938^age
**Male (Cr <= 0.9):** 142 x (Cr/0.9)^(-0.302) x 0.9938^age
**Male (Cr > 0.9):** 142 x (Cr/0.9)^(-1.200) x 0.9938^age

## AKI Staging (KDIGO)

| Stage | Creatinine Criteria | Urine Output |
|-------|-------------------|-------------|
| 1 | Increase >= 0.3 mg/dL within 48h OR 1.5-1.9x baseline within 7 days | <0.5 mL/kg/h for 6-12h |
| 2 | 2.0-2.9x baseline | <0.5 mL/kg/h for >= 12h |
| 3 | 3.0x baseline OR Cr >= 4.0 OR initiation of RRT | <0.3 mL/kg/h for >= 24h OR anuria >= 12h |

## Medication Dose Adjustments by eGFR

| Medication | eGFR >60 | eGFR 30-59 | eGFR 15-29 | eGFR <15 |
|-----------|----------|-----------|-----------|----------|
| Metformin | Full dose | Reduce to 1000mg/d (eGFR 30-45); caution 45-60 | Contraindicated | Contraindicated |
| Gabapentin | Up to 3600mg/d | 300-700mg/d | 100-300mg/d | 100-300mg/d |
| Pregabalin | Up to 600mg/d | Up to 300mg/d | Up to 150mg/d | 25-75mg/d |
| Apixaban | 5mg BID | 5mg BID (unless 2 of: age>80, wt<60, Cr>1.5) | 5mg BID (limited data) | Not recommended |
| Rivaroxaban | 20mg daily | 15mg daily (eGFR 15-50) | Avoid (eGFR <15) | Avoid |
| Dabigatran | 150mg BID | 150mg BID (eGFR >30) | Avoid (eGFR <30) | Avoid |
| Enoxaparin | 1mg/kg BID | 1mg/kg BID | 1mg/kg DAILY | Avoid |
| Allopurinol | Up to 800mg/d | 200mg/d | 100mg/d | 100mg QOD |
| Vancomycin | Per levels | Per levels, extend interval | Per levels, extend interval | Per levels only |
| Acyclovir | Full dose | Reduce dose/interval | Reduce significantly | 50% dose, extend interval |

## Nephrology Referral Criteria

Refer to nephrology when:
- eGFR <30 mL/min/1.73m2 (CKD G4-G5)
- Rapid eGFR decline: >5 mL/min/1.73m2 per year
- Persistent albuminuria: UACR >300 mg/g
- Refractory hypertension (>3 agents at max dose)
- Persistent electrolyte abnormalities (hyperkalemia, metabolic acidosis)
- Recurrent nephrolithiasis
- Hereditary kidney disease
- Unexplained hematuria with proteinuria
