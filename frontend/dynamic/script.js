let clicked = 0;
const dropdown = document.querySelectorAll('.dropdown i')
let contentList = document.querySelectorAll('.content')

dropdown.forEach(drop => {
    
    let contentSibling = drop.nextElementSibling
    drop.addEventListener('click', ()=>{
        contentSibling.classList.toggle("show")
    })
    document.addEventListener('click', function(event) {
        if (!contentSibling.contains(event.target) && !drop.contains(event.target) && contentSibling.classList.contains("show")) {
            console.log(contentSibling.classList);
            contentSibling.classList.remove('show');
        }
    });
})

function checkUserIsLogged(){
    localStorage.setItem("Token", "This is a test token")
    const token = localStorage.getItem("token")
}

// for Comments Like and Dislike on Post Page
let like = document.querySelectorAll(".react .like")
let dislike = document.querySelectorAll(".react .dislike")

// Handling Like Button Clicked in Post Comments
like.forEach(like_elem => {
    like_elem.addEventListener('click', function(){
        // Check if the user is loggedin or not :
        // And check if Dislike is already clicked
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

let allContents = document.querySelectorAll(".tweet-text")

allContents.forEach(content => {
    content.addEventListener('click', function(event){
        if (content.classList.contains("collapse"))
            content.classList.remove("collapse")
        else
            content.classList.add("collapse")

    })
})