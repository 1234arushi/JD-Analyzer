//request -> message, sender->popup/background , sendResponse-> function to send data back
chrome.runtime.onMessage.addListener((request,sender,sendResponse)=>{
    if(request.action==="extractJD"){
        console.log("Extracting Job Description..")
        const companyElement=document.querySelector(
            //find first <a> tag that lives inside any element & whose aria-label starts with "Company"
            '[aria-label^="Company"] a'// ^= -> starts-with operator
            //aria-label : attribute(for describing element) must be enclosed in sq bracket
            // a -> aria-label must have a link inside
        )
        const lazyColumn = document.querySelector(
            '[data-testid="lazy-column"]'
        );
        const paragraphs = lazyColumn.querySelectorAll("p");
       
        const roleElement = paragraphs[1];

        const locationElement = paragraphs[2];

        // Job Description
        const jdElement = document.querySelector(
            '[data-testid="expandable-text-box"]'
        );
        //sending data back to popup.js

        sendResponse({
            company:companyElement//ternary operation
                   ?companyElement.textContent.trim()
                   :"",
            role:roleElement
                    ? roleElement.textContent.trim()
                    : "",
            location:locationElement
                    ?locationElement.textContent.trim()
                    :"",
            jobDescription:jdElement
                    ?jdElement.textContent.trim()
                    :"",
                
        });
    }

});