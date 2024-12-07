// for Favourite Button on post Page
let favourite_area = document.querySelector(".addToFavourite")
let favourite_icon = document.querySelector(".addToFavourite span")

// for Comments Like and Dislike on Post Page
let comment_like = document.querySelector(".commentReaction .like")
let comment_dislike = document.querySelector(".commentReaction .dislike")

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
    }else{
        favourite_icon.textContent = "bookmark_add"
        favourite_icon.style.color = "#919191"

    }
})
// Handling Like Button Clicked in Post Comments
comment_like.addEventListener('click', function(){
    if (comment_dislike.classList.contains("clicked"))
        comment_dislike.classList.remove("clicked");
    comment_like.classList.toggle("clicked");
})
comment_dislike.addEventListener('click', function(){
    if (comment_like.classList.contains("clicked"))
        comment_like.classList.remove("clicked");
    comment_dislike.classList.toggle("clicked");
})