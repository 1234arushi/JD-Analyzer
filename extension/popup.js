const analyzeBtn = document.getElementById("analyze-button");
const status=document.getElementById("status");

analyzeBtn.addEventListener("click",()=>{
    status.textContent="Analyzing...";
    chrome.tabs.query(
        {active:true,currentWinow:true},
        //after receiving a tab, we are asking chrome to go to content.js for that
        (tabs)=>{
            chrome.tabs.sendMessage(
                tabs[0].id,
                {action:"extractJD"},
                async (response)=>{
                    status.textContent="Extraction Complete!";
                    const res=await fetch("http://localhost:8080/jdscope/v1/analyze",{
                        method:"POST",
                        headers:{
                            "Content-type":"application/json",
                        },
                        body:JSON.stringify(response),
                    });
                    const data=await res.json();
                }
            )
        }
    )

});

