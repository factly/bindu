let binduChartsURL = "http://localhost:7002/charts/"
let embedElements = document.getElementsByClassName("factly-embed");

for(let i=0; i<embedElements.length; i++) {
        let chartID = embedElements[i].getAttribute("data-src");
        fetch(binduChartsURL+chartID).then((res) => {
                return res.json();
        }).then((json) => {
                embedElements[i].setAttribute('id', 'embed'+i)
                vegaEmbed('#embed'+i, json).catch(console.error);
        }).catch((err) => {
                console.log(err);
        })
}