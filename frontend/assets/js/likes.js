
// Handling Likes and Dislikes in both the frontend and backend

function handleLikes() {
    const likeBtns = document.querySelectorAll('.like');
    const dislikeBtns = document.querySelectorAll('.dislike');

    likeBtns.forEach(btn => {
        btn.addEventListener('click', async () => {
            console.clear();
            const postId = btn.id;
            const postOrComment = (btn.getAttribute('isPost') == "true")? true : false;
            
            console.log("Post or Comment = ", postOrComment, btn);
            
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
                console.log("Response = ", res);
                console.log("Response status = ", res.status);
                
                if (res.status == 401){
                    popUp()
                    console.log("Error");
                    return;
                }else if (res.status != 200){
                    console.log("Error response status = ", res.status);
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
            console.clear();

            const postId = btn.id;
            const postOrComment = (btn.getAttribute('isPost') == "true")? true : false;
            console.log("Post or Comment = ", postOrComment, btn);

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
                console.log(res);
                console.log(res.status);
                
                if (res.status == 401){
                    popUp()
                    console.log("Error");
                    return;
                }else if (res.status != 200){
                    console.log("Error response status = ", res.status);
                    return;
                }
                const data = await res.json();
                console.log(data);
                const like = btn.previousElementSibling
                
                if (data.added){
                    btn.classList.add("FILL")
                    if (like) like.classList.remove("FILL")
                }else{
                    btn.classList.remove("FILL")
                }
                if (like) like.querySelector('span').innerText = data.likes;
                btn.querySelector('span').innerText = data.dislikes;
            
            }catch(error){
                console.log(error)
            }   
         });
    });
}

handleLikes();
