let embedElements = document.getElementsByClassName("factly-embed");

for(let i=0; i<embedElements.length; i++) {
        let embedSrc = embedElements[i].getAttribute("data-src");
        embedElements[i].setAttribute('id', 'embed'+i)
        
        vegaEmbed('#embed'+i, embedSrc).catch(console.error);
}