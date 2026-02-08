# Mental Health Screening Tools Reference

## PHQ-2 (Patient Health Questionnaire-2)

### Purpose
Brief depression screener. First 2 items of PHQ-9. Used as initial screen; if positive, administer full PHQ-9.

### Items

| # | Question | Score 0-3 |
|---|----------|-----------|
| 1 | Over the last 2 weeks, how often have you been bothered by little interest or pleasure in doing things? | 0=Not at all, 1=Several days, 2=More than half the days, 3=Nearly every day |
| 2 | Over the last 2 weeks, how often have you been bothered by feeling down, depressed, or hopeless? | 0=Not at all, 1=Several days, 2=More than half the days, 3=Nearly every day |

### Scoring
- Range: 0-6
- Positive screen: >= 3
- LOINC: 69737-5

### Interpretation and Action

| Score | Interpretation | Action |
|-------|---------------|--------|
| 0-2 | Negative screen | No further depression screening needed at this visit |
| 3-6 | Positive screen | Administer full PHQ-9 |

### Performance
- Sensitivity: 83% for major depression (score >= 3)
- Specificity: 92%

## PHQ-9 (Patient Health Questionnaire-9)

### Purpose
Depression severity assessment. Validated for screening and monitoring treatment response.

### Items

Over the last 2 weeks, how often have you been bothered by:

| # | Question | LOINC | Score 0-3 |
|---|----------|-------|-----------|
| 1 | Little interest or pleasure in doing things | 44250-9 | 0-3 |
| 2 | Feeling down, depressed, or hopeless | 44255-8 | 0-3 |
| 3 | Trouble falling or staying asleep, or sleeping too much | 44259-0 | 0-3 |
| 4 | Feeling tired or having little energy | 44254-1 | 0-3 |
| 5 | Poor appetite or overeating | 44251-7 | 0-3 |
| 6 | Feeling bad about yourself, or that you are a failure, or have let yourself or family down | 44258-2 | 0-3 |
| 7 | Trouble concentrating on things, such as reading the newspaper or watching television | 44252-5 | 0-3 |
| 8 | Moving or speaking so slowly that other people could have noticed? Or the opposite, being so fidgety or restless that you have been moving around a lot more than usual | 44253-3 | 0-3 |
| 9 | Thoughts that you would be better off dead or of hurting yourself in some way | 44261-6 | 0-3 |

Scoring: 0=Not at all, 1=Several days, 2=More than half the days, 3=Nearly every day

### Scoring and Severity

- Range: 0-27
- LOINC (total score): 44249-1

| Score | Severity | Recommended Action |
|-------|----------|-------------------|
| 0-4 | Minimal / None | No treatment needed. Rescreen per schedule. |
| 5-9 | Mild | Watchful waiting. Rescreen in 4 weeks. Consider counseling if persistent. |
| 10-14 | Moderate | Treatment plan: psychotherapy (CBT, IPT), consider pharmacotherapy. Follow up in 2-4 weeks. |
| 15-19 | Moderately severe | Active treatment: pharmacotherapy AND/OR psychotherapy. Close follow-up in 1-2 weeks. |
| 20-27 | Severe | Immediate pharmacotherapy + psychotherapy. Consider psychiatric referral. Assess safety (item 9). Follow up in 1 week. |

### Item 9 (Suicidal Ideation) Special Handling

| Item 9 Score | Meaning | Required Action |
|-------------|---------|----------------|
| 0 | No suicidal ideation | Document negative screen |
| 1 (Several days) | Intermittent passive SI | Administer C-SSRS. Safety assessment. Outpatient mental health referral within 1 week. |
| 2 (More than half the days) | Frequent passive/active SI | Administer C-SSRS immediately. Same-day mental health evaluation. Safety plan. Lethal means counseling. |
| 3 (Nearly every day) | Daily suicidal thoughts | Administer C-SSRS immediately. Do not leave patient unattended. Emergency psychiatric evaluation. Assess need for hospitalization. |

### Treatment Response Monitoring

- Clinically significant change: >= 5-point decrease in PHQ-9 score
- Remission target: PHQ-9 < 5
- Partial response: 25-49% reduction in score
- Response: >= 50% reduction in score
- Rescreen interval during treatment: every 2-4 weeks until stable, then every 3 months

## GAD-7 (Generalized Anxiety Disorder-7)

### Purpose
Anxiety severity screening and monitoring. Validated for GAD; also screens for panic disorder, social anxiety, and PTSD (with lower sensitivity).

### Items

Over the last 2 weeks, how often have you been bothered by:

| # | Question | LOINC | Score 0-3 |
|---|----------|-------|-----------|
| 1 | Feeling nervous, anxious, or on edge | 69725-0 | 0-3 |
| 2 | Not being able to stop or control worrying | 68509-9 | 0-3 |
| 3 | Worrying too much about different things | 69733-4 | 0-3 |
| 4 | Trouble relaxing | 69734-2 | 0-3 |
| 5 | Being so restless that it's hard to sit still | 69735-9 | 0-3 |
| 6 | Becoming easily annoyed or irritable | 69689-8 | 0-3 |
| 7 | Feeling afraid as if something awful might happen | 69736-7 | 0-3 |

Scoring: 0=Not at all, 1=Several days, 2=More than half the days, 3=Nearly every day

### Scoring and Severity

- Range: 0-21
- LOINC (total score): 75626-2

| Score | Severity | Recommended Action |
|-------|----------|-------------------|
| 0-4 | Minimal | No treatment needed. Rescreen per schedule. |
| 5-9 | Mild | Watchful waiting. Psychoeducation. Relaxation techniques. Rescreen in 4 weeks. |
| 10-14 | Moderate | Treatment plan: CBT (first-line) and/or SSRI/SNRI. Follow up in 2-4 weeks. |
| 15-21 | Severe | Active treatment: pharmacotherapy + psychotherapy. Consider psychiatric referral. Follow up in 1-2 weeks. |

### Performance
- Sensitivity: 89% for GAD (score >= 10)
- Specificity: 82%
- Also screens for: panic disorder (sensitivity 74%), social anxiety (sensitivity 72%), PTSD (sensitivity 66%)

## AUDIT-C (Alcohol Use Disorders Identification Test - Consumption)

### Purpose
Brief alcohol use screening. First 3 items of the full 10-item AUDIT.

### Items

| # | Question | Scoring |
|---|----------|---------|
| 1 | How often do you have a drink containing alcohol? | 0=Never, 1=Monthly or less, 2=2-4x/month, 3=2-3x/week, 4=4+x/week |
| 2 | How many drinks containing alcohol do you have on a typical day when you are drinking? | 0=1-2, 1=3-4, 2=5-6, 3=7-9, 4=10+ |
| 3 | How often do you have 6 or more drinks on one occasion? | 0=Never, 1=Less than monthly, 2=Monthly, 3=Weekly, 4=Daily or almost daily |

### Scoring and Interpretation

- Range: 0-12
- LOINC: 77564-3

| Population | Positive Screen | Action |
|-----------|----------------|--------|
| Men | >= 4 | Further assessment with full AUDIT or clinical evaluation. Brief intervention. |
| Women | >= 3 | Further assessment with full AUDIT or clinical evaluation. Brief intervention. |
| Any | Score of 0 with prior history | Verify abstinence is intentional. Assess recovery status. |

### Risk Stratification

| AUDIT-C Score | Risk Level |
|--------------|-----------|
| 0 | Non-drinker (or in recovery) |
| 1-3 (men) / 1-2 (women) | Low-risk drinking |
| 4-7 (men) / 3-7 (women) | Hazardous drinking -- brief intervention |
| 8-12 | Likely alcohol use disorder -- referral for evaluation |

## Columbia Suicide Severity Rating Scale (C-SSRS)

### Purpose
Structured suicide risk assessment. Measures suicidal ideation severity and behavior. FDA-endorsed for clinical trials and clinical practice.

### Suicidal Ideation Severity Scale

| Level | Category | Description | LOINC |
|-------|----------|-------------|-------|
| 0 | No ideation | Denies suicidal thoughts | -- |
| 1 | Wish to be dead | "I wish I were dead" or "I wish I could go to sleep and not wake up" | 89206-7 |
| 2 | Non-specific active suicidal thoughts | General thoughts of wanting to end one's life without specific method, plan, or intent | 89206-7 |
| 3 | Active suicidal ideation with any methods (not plan) without intent to act | Thinking of a method but no specific plan or intent | 89206-7 |
| 4 | Active suicidal ideation with some intent to act, without specific plan | Some intent to act on suicidal thoughts but no specific plan | 89206-7 |
| 5 | Active suicidal ideation with specific plan and intent | Has worked out details of a plan and intends to carry it out | 89206-7 |

### Suicidal Behavior Categories

| Category | Description |
|----------|-------------|
| Actual attempt | Self-injurious behavior with intent to die (regardless of injury) |
| Interrupted attempt | Interrupted by self or outside circumstance before harm |
| Aborted attempt | Started to take steps but stopped before actual attempt |
| Preparatory acts | Collecting pills, buying weapon, writing note, giving away possessions |
| Non-suicidal self-injury | Self-harm without intent to die |

### Risk Triage by C-SSRS Level

| C-SSRS Level | Risk Category | Disposition |
|-------------|--------------|-------------|
| 0 | No identified risk | Standard care. Document negative screen. |
| 1-2 | Low risk | Outpatient mental health referral. Safety plan. Lethal means counseling. Follow up 1 week. |
| 3 | Moderate risk | Same-day mental health evaluation. Safety plan. Lethal means restriction. Consider whether safe for discharge. |
| 4-5 | High risk | Emergency psychiatric evaluation. Do not leave unattended. Consider inpatient admission. Contact crisis team. |
| Actual/interrupted attempt | Acute risk | Emergency department. Medical stabilization. Psychiatric hold as indicated. |

## Mood Disorder Questionnaire (MDQ)

### Purpose
Bipolar disorder screening. Administer before initiating antidepressant therapy in patients with mood symptoms, especially if history of mood cycling, irritability, or family history of bipolar disorder.

### Items

**Part 1 (13 yes/no items):** Has there ever been a period of time when you were not your usual self and...

| # | Item |
|---|------|
| 1 | You felt so good or so hyper that other people thought you were not your normal self, or you were so hyper that you got into trouble? |
| 2 | You were so irritable that you shouted at people or started fights or arguments? |
| 3 | You felt much more self-confident than usual? |
| 4 | You got much less sleep than usual and found you didn't really miss it? |
| 5 | You were much more talkative or spoke much faster than usual? |
| 6 | Thoughts raced through your head or you couldn't slow your mind down? |
| 7 | You were so easily distracted by things around you that you had trouble concentrating or staying on track? |
| 8 | You had much more energy than usual? |
| 9 | You were much more active or did many more things than usual? |
| 10 | You were much more social or outgoing than usual; for example, you telephoned friends in the middle of the night? |
| 11 | You were much more interested in sex than usual? |
| 12 | You did things that were unusual for you or that other people might have thought were excessive, foolish, or risky? |
| 13 | Spending money got you or your family into trouble? |

**Part 2:** If you checked YES to more than one of the above, have several of these ever happened during the same period of time?

**Part 3:** How much of a problem did any of these cause you -- like being unable to work; having family, money, or legal troubles; getting into arguments or fights? (No problem / Minor problem / Moderate problem / Serious problem)

### Scoring

Positive screen requires ALL three:
1. >= 7 "yes" answers in Part 1
2. "Yes" to Part 2 (co-occurrence)
3. "Moderate" or "Serious" problem in Part 3

### Performance
- Sensitivity: 73% for bipolar I and II
- Specificity: 90%

### Clinical Implications of Positive MDQ
- Do NOT start antidepressant monotherapy (risk of mania switch)
- Refer to psychiatry for diagnostic evaluation
- Consider mood stabilizer if treatment needed urgently (lithium, valproate, or quetiapine)

## PC-PTSD-5 (Primary Care PTSD Screen for DSM-5)

### Purpose
Brief PTSD screening for primary care. 5 yes/no items.

### Stem Question
"Sometimes things happen to people that are unusually or especially frightening, horrible, or traumatic. For example: a serious accident or fire, a physical or sexual assault or abuse, an earthquake or flood, a war, seeing someone be killed or seriously injured, having a loved one die through homicide or suicide."

"Have you ever experienced this kind of event?" If YES, proceed:

### Items

| # | Question | Score |
|---|----------|-------|
| 1 | Had nightmares about the event(s) or thought about the event(s) when you did not want to? | 0=No, 1=Yes |
| 2 | Tried hard not to think about the event(s) or went out of your way to avoid situations that reminded you of the event(s)? | 0=No, 1=Yes |
| 3 | Been constantly on guard, watchful, or easily startled? | 0=No, 1=Yes |
| 4 | Felt numb or detached from people, activities, or your surroundings? | 0=No, 1=Yes |
| 5 | Felt guilty or unable to stop blaming yourself or others for the event(s) or any problems the event(s) may have caused? | 0=No, 1=Yes |

### Scoring and Interpretation

- Range: 0-5
- LOINC: not yet assigned standard code; use local code or generic survey Observation

| Score | Interpretation | Action |
|-------|---------------|--------|
| 0-2 | Negative screen | No further PTSD evaluation needed at this visit |
| 3-5 | Positive screen | Refer for comprehensive PTSD evaluation (PCL-5 or clinical interview). Consider trauma-focused therapy referral. |

### Performance (cutoff >= 3)
- Sensitivity: 95%
- Specificity: 85%

## Screening Frequency Recommendations

| Tool | Population | Frequency |
|------|-----------|-----------|
| PHQ-2/PHQ-9 | All adults >= 18 (USPSTF) | Annual at minimum; every visit if risk factors; every 2-4 weeks during treatment |
| GAD-7 | Adults with anxiety symptoms or risk factors | At initial assessment; every 2-4 weeks during treatment; annual screening reasonable |
| AUDIT-C | All adults (USPSTF) | Annual |
| C-SSRS | Any positive suicidality screen | Immediately when triggered; at every visit if prior positive |
| MDQ | Adults with depressive symptoms before antidepressant initiation | Once (at initial evaluation) |
| PC-PTSD-5 | Adults with trauma history or symptoms | Annual in primary care; at intake for behavioral health |
