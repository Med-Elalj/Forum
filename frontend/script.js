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


// for Comments Like and Dislike on Post Page
let like = document.querySelectorAll(".react .like")
let dislike = document.querySelectorAll(".react .dislike")

// Handling Like Button Clicked in Post Comments
like.forEach(like_elem => {
    like_elem.addEventListener('click', function(){
        let dislike_sibling = like_elem.nextElementSibling;
        like_elem.classList.toggle("FILL");
        dislike_sibling.classList.remove("FILL");
    })
})

dislike.forEach(dislike_elem => {
    dislike_elem.addEventListener('click', function(){
        let like_sibling = dislike_elem.previousElementSibling;
        dislike_elem.classList.toggle("FILL");
        like_sibling.classList.remove("FILL");

        // the  rest of the code will write here to send request
        // to backend to update database 
        
    })
})