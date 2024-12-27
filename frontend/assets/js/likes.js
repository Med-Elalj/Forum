
// Handling Likes and Dislikes in both the frontend and backend

function handleLikes() {
    const likeBtns = document.querySelectorAll('.like');
    const dislikeBtns = document.querySelectorAll('.dislike');

    likeBtns.forEach(btn => {
        btn.addEventListener('click', async () => {
            const postId = btn.id;
            const postOrComment = (btn.getAttribute('isPost') == "true")? true : false;
            console.log("Post or Comment = ", postOrComment);
            
            try {
                const res = await fetch(`/PostReaction`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ 
                        postId: postId,
                        "type": "like",
                        post: postOrComment
                     })
                });
                console.log("postOrComment", postOrComment);
                console.log("Response = ", res);
                
                if (res.status == 401){
                    popUp()
                    console.log("Error");
                    return;
                }
                const data = await res.json();
                console.log(data);
                
                const dislike = btn.nextElementSibling
                if (data.added){
                    btn.classList.add("FILL")
                    dislike.classList.remove("FILL")
                }else{
                    btn.classList.remove("FILL")
                }
                dislike.querySelector('span').innerText = data.dislikes;
                btn.querySelector('span').innerText = data.likes;
            } catch (error) {
                console.log(error)
            }
          
        });
    });

    dislikeBtns.forEach(btn => {
        btn.addEventListener('click', async () => {
            const postId = btn.id;
            const postOrComment = (btn.getAttribute('isPost') == "true")? true : false;
            console.log("Post or Comment = ", postOrComment);

            try{
                const res = await fetch(`/PostReaction`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ 
                        postId: postId,
                        "type": "dislike",
                        post: postOrComment

                     })
                });
    
                if (res.status == 401){
                    popUp()
                    console.log("Error");
                    return;
                }
                const data = await res.json();
                console.log(data);
                const like = btn.previousElementSibling
                console.log("=====> Added = ", data.added);
                
                if (data.added){
                    btn.classList.add("FILL")
                    like.classList.remove("FILL")
                }else{
                    btn.classList.remove("FILL")
                }
                like.querySelector('span').innerText = data.likes;
                btn.querySelector('span').innerText = data.dislikes;
            
            }catch(error){
                console.log(error)
            }   
         });
    });
}

handleLikes();
