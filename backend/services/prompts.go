package services

const CleanJDPrompt = `You are given a scraped job description from a job portal.

Remove company marketing, culture descriptions, legal notices, diversity statements, hiring process, navigation text, and any unrelated content.

Preserve only information relevant to the actual job, including responsibilities, requirements, qualifications, skills, experience, education and role details.

Do not summarize.

Return only the cleaned job description.

Job Description:

`
