
// Handling Likes and Dislikes in both the frontend and backend

async function LikePostAndComments(btn){
    
        const postId = btn.id;
        const postOrComment = (btn.getAttribute('isPost') == "true")? true : false;
        
        
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
            if (res.status == 401){
                popUp()
                return;
            }else if (res.status != 200){
                return;
            }
            const data = await res.json();
            
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
            errorr = error
        }
      
}

async function DisLikePostAndComments(btn){

    const postId = btn.id;
    const postOrComment = (btn.getAttribute('isPost') == "true")? true : false;

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
            return;
        }else if (res.status != 200){
            return;
        }
        const data = await res.json();
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
        errorr = error
    }   
}
function handleLikes(add) {

    const likeBtns = document.querySelectorAll('.like');
    const dislikeBtns = document.querySelectorAll('.dislike');

    likeBtns.forEach(btn => {
        if (add){
            btn.addEventListener('click', () => LikePostAndComments(btn));
        }else{
            btn.removeEventListener('click', () => LikePostAndComments(btn));
        }
    });

    dislikeBtns.forEach(btn => {
        if (add){
            btn.addEventListener('click', () => DisLikePostAndComments(btn));
        }else{
            btn.removeEventListener('click', () => DisLikePostAndComments(btn));
        }
    });
}