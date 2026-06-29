package services

const CleanJDPrompt = `You are given a scraped job description.

Extract ONLY the following information.

Return STRICT JSON.

{
  "required_skills": [],
  "nice_to_have_skills": [],
  "experience": ""
}

Rules:

- Ignore company culture, benefits, diversity statements, hiring process and marketing.
- Include only explicitly mentioned skills.
- Put mandatory skills in required_skills.
- Put preferred/additional skills which are less imp according to job poster in nice_to_have_skills,otherwise leave it empty.
- Do not return explanations,simply keywords.
- Return valid JSON only.
- Do not wrap the response in json tags
- Do not use markdown.
- Do not add explanations.
- Do not add any text before or after the JSON.

`
