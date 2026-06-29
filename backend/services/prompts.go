package services

const CleanJDPrompt = `You are given a scraped job description.

Extract ONLY the following information.

Return STRICT JSON.

{
  "required_skills": [],
  "nice_to_have_skills": [],
  "education": [],
  "experience": ""
}

Rules:

- Ignore company culture, benefits, diversity statements, hiring process and marketing.
- Include only explicitly mentioned skills.
- Put mandatory skills in required_skills.
- Put preferred skills in nice_to_have_skills.
- Do not return explanations.
- Return valid JSON only.

`
