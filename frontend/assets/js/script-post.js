function x1(){  


// for Favourite Button on post Page
// let favourite_area = document.querySelector(".addToFavourite")
// let favourite_icon = document.querySelector(".addToFavourite span")

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

// favourite_area.addEventListener('click', function(){
//     if (favourite_icon.textContent == "bookmark_add"){
//         favourite_icon.textContent = "bookmark_added"
//         favourite_icon.style.color = "#088395"
//         favourite_icon.classList.add("FILL")
//          // the  rest of the code will write here to send request
//         // to backend to update database 
//     }else{
//         favourite_icon.textContent = "bookmark_add"
//         favourite_icon.style.color = "#919191"
//         favourite_icon.classList.remove("FILL")

//          // the  rest of the code will write here to send request
//         // to backend to update database 

//     }
// })

// for Comments Like and Dislike on Post Page
let comment_like = document.querySelectorAll(".commentReaction .like")
let comment_dislike = document.querySelectorAll(".commentReaction .dislike")

// Handling Like Button Clicked in Post Comments
comment_like.forEach(like => {
    like.addEventListener('click', function(){
        if (checkUserIsLogged()){
            // Check from DB if like or dislike exits
            /// /// / // / / / / / /
            let dislike = like.nextElementSibling;
            like.classList.toggle("FILL");
            dislike.classList.remove("FILL");
        }else{
            popUp();
        }
       
    })
})

comment_dislike.forEach(dislike => {
    dislike.addEventListener('click', function(){
        let like = dislike.previousElementSibling;
        dislike.classList.toggle("FILL");
        like.classList.remove("FILL");

        // the  rest of the code will write here to send request
        // to backend to update database 
        
    })
})
}

function CommentReactionEventListenner(){
    const commentButtons = document.querySelectorAll('.commentReaction .like');
    commentButtons.forEach(button => {
        button.addEventListener('click', async (e) => {
            const commentID = e.target.id;
            const response = await fetch(`/LikeComment/${commentID}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                }
            });
            const data = await response.json();
            if (data.success) {
                const likeCount = e.target.querySelector('span');
                likeCount.innerText = data.likeCount;// [ 'likeCount' : 1, 'dislikeCount' : 0 ];
            }
        });
    });
    const dislikeButtons = document.querySelectorAll('.commentReaction .dislike');
    dislikeButtons.forEach(button => {
        button.addEventListener('click', async (e) => {
            const commentID = e.target.id;
            const response = await fetch(`/DislikeComment/${commentID}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                }
            });
            const data = await response.json();
            if (data.success) {
                const dislikeCount = e.target.querySelector('span');
                dislikeCount.innerText = data.dislikeCount;
            }
        });
    });
}

function CommentInputEventListenner(){
    console.log("Listx");

    const commentInput = document.querySelector('.commentInput input');
    console.log("List");
    commentInput.addEventListener('keypress', async (e) => {
        
        console.log("Rak Hena ");
        
        if (e.key === 'Enter') {
            console.log('Enter Pressed');
            e.preventDefault();
            const comment = commentInput.value;
            commentInput.value = '';
            const postID = commentInput.id;
            const response = await fetch('/CreateComment', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    postID,
                    comment
                })
            });
            console.log("response", response);
            
            const data = await response.json();
            console.log("Response  = data", data);
            console.log("data stst", data["status"] == "ok");
            console.log('Comment Added', (response.success));
            if (data["status"] == "ok") {
                const commentContainer = document.querySelector('.Comments');
                const commentCard = document.createElement('div');
                commentCard.classList.add('commentCard');
                commentCard.classList.add('CommentAdded');
                commentCard.id = "Comment" + data["CommentID"];
                console.log(commentCard);
                let url = "/index#Comment" + data["CommentID"];
                commentCard.innerHTML = `
                    <div class="commentAuthorImage">
                        <img src="https://ui-avatars.com/api/?name=${data["UserName"]}" alt="">
                    </div>
                    <div class="commentAuthor">
                        <div class="commentAuthorInfo">
                            <span class="commentAuthorName">
                                @${data["UserName"]}
                                <span class="commentTime">
                                    ${data["CreatedAt"]}
                                </span>
                            </span>
                            <div class="commentReaction DisableUserSelect">
                                <span class="like" id="${data["CommentID"]}">
                                    <span class="material-symbols-outlined">
                                        thumb_up
                                    </span>
                                    <span>0</span>
                                </span>
                                <span class="dislike" id="${data["CommentID"]}">
                                    <span class="material-symbols-outlined">thumb_down</span>
                                    <span>0</span>
                                </span>
                            </div>
                        </div>
                        <div class="commentInfo">
                            <p class="commentData collapse">${data["Content"]}</p>
                        </div>
                    </div>
                `;
                commentContainer.prepend(commentCard);
                window.location.replace(url)
                CommentReactionEventListenner(); //Re-adding event listeners
            }
        }
    });
}
CommentReactionEventListenner()
CommentInputEventListenner()
x1()