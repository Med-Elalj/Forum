console.log("TEST");
let clicked = 0;
const dropdown = document.querySelector('.dropdown i')
const contentList = document.querySelector('.content')
dropdown.addEventListener('click', ()=>{
    contentList.classList.toggle("show")
})

document.addEventListener('click', function(event) {
    if (!contentList.contains(event.target) && !dropdown.contains(event.target) && contentList.classList.contains("show")) {
        console.log(contentList.classList);
        contentList.classList.remove('show');
    }
});