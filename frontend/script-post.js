// for Favourite Button on post Page
let favourite_area = document.querySelector(".addToFavourite")
let favourite_icon = document.querySelector(".addToFavourite span")

// For read More in Post Page
let comments = document.querySelectorAll(".commentData")

comments.forEach(elem => {
    elem.addEventListener('click', ()=>{
        elem.classList.toggle("collapse")
        comments.forEach(second_elem => {
            if (second_elem != elem)
                second_elem.classList.add("collapse")
        })
    })
})

favourite_area.addEventListener('click', function(){
    if (favourite_icon.textContent == "bookmark_add"){
        favourite_icon.textContent = "bookmark_added"
        favourite_icon.style.color = "#088395"
         // the  rest of the code will write here to send request
        // to backend to update database 
    }else{
        favourite_icon.textContent = "bookmark_add"
        favourite_icon.style.color = "#919191"
         // the  rest of the code will write here to send request
        // to backend to update database 

    }
})

// for Comments Like and Dislike on Post Page
let comment_like = document.querySelectorAll(".commentReaction .like")
let comment_dislike = document.querySelectorAll(".commentReaction .dislike")

// Handling Like Button Clicked in Post Comments
comment_like.forEach(like => {
    like.addEventListener('click', function(){
        let dislike = like.nextElementSibling;
        if (dislike.classList.contains("clicked"))
            dislike.classList.remove("clicked");
        like.classList.toggle("clicked");
         // the  rest of the code will write here to send request
        // to backend to update database 
    })
})

comment_dislike.forEach(dislike => {
    dislike.addEventListener('click', function(){
        let like = dislike.previousElementSibling;
        if (like.classList.contains("clicked"))
            like.classList.remove("clicked");
        dislike.classList.toggle("clicked");
        // the  rest of the code will write here to send request
        // to backend to update database 
        
    })
})