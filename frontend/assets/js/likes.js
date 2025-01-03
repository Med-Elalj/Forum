// Store handler references
const likeHandlers = new WeakMap();
const dislikeHandlers = new WeakMap();

async function LikePostAndComments(btn) {
    const postId = btn.id;
    const postOrComment = btn.getAttribute('isPost') === "true";
    console.log(postOrComment);
    
    try {
        const res = await fetch('/PostReaction', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ 
                postId,
                type: "like",
                post: postOrComment
            })
        });

        if (res.status === 401) {
            popUp();
            return;
        }
        if (res.status !== 200) {
            return;
        }

        const data = await res.json();
        const dislike = btn.nextElementSibling;
        
        if (data.added) {
            btn.classList.add("FILL");
            dislike.classList.remove("FILL");
        } else {
            btn.classList.remove("FILL");
        }
        
        dislike.querySelector('span').innerText = data.dislikes;
        btn.querySelector('span').innerText = data.likes;
    } catch (error) {
        console.error('Like operation failed:', error);
    }
}

async function DisLikePostAndComments(btn) {
    const postId = btn.id;
    const postOrComment = btn.getAttribute('isPost') === "true";
    console.log(postOrComment);

    try {
        const res = await fetch('/PostReaction', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ 
                postId,
                type: "dislike",
                post: postOrComment
            })
        });
        
        if (res.status === 401) {
            popUp();
            return;
        }
        if (res.status !== 200) {
            return;
        }

        const data = await res.json();
        const like = btn.previousElementSibling;
        
        if (data.added) {
            btn.classList.add("FILL");
            if (like) like.classList.remove("FILL");
        } else {
            btn.classList.remove("FILL");
        }
        
        if (like) like.querySelector('span').innerText = data.likes;
        btn.querySelector('span').innerText = data.dislikes;
    } catch (error) {
        console.error('Dislike operation failed:', error);
    }
}

function handleLikes(add) {
    const likeBtns = document.querySelectorAll('.like');
    const dislikeBtns = document.querySelectorAll('.dislike');

    likeBtns.forEach(btn => {
        // Remove old handler if exists
        if (likeHandlers.has(btn)) {
            btn.removeEventListener('click', likeHandlers.get(btn));
        }
        
        if (add) {
            // Create and store new handler
            const handler = () => LikePostAndComments(btn);
            likeHandlers.set(btn, handler);
            btn.addEventListener('click', handler);
        }
    });

    dislikeBtns.forEach(btn => {
        // Remove old handler if exists
        if (dislikeHandlers.has(btn)) {
            btn.removeEventListener('click', dislikeHandlers.get(btn));
        }
        
        if (add) {
            // Create and store new handler
            const handler = () => DisLikePostAndComments(btn);
            dislikeHandlers.set(btn, handler);
            btn.addEventListener('click', handler);
        }
    });
}