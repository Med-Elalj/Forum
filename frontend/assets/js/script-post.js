function ExpandComments() {
    // Expand Comment and read Content...
    let comments = document.querySelectorAll(".commentData")

    comments.forEach(elem => {
        elem.addEventListener('click', () => {
            elem.classList.toggle("collapse")
            comments.forEach(second_elem => {
                if (second_elem != elem)
                    second_elem.classList.add("collapse")
            })
        })
    })
}

function CommentInputEventListenner() {
    const commentInput = document.querySelector('.commentInput input');

    commentInput.addEventListener('keypress', async (e) => {
        if (e.key === 'Enter') {
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
            if (response.status != 200) {
                console.log("Error : Inside Response Status");
                popUp()
                return
            }

            const data = await response.json();
            if (data["status"] == "ok") {
                const commentContainer = document.querySelector('.Comments');
                const commentCard = document.createElement('div');
                commentCard.classList.add('commentCard');
                commentCard.classList.add('CommentAdded');
                commentCard.id = "Comment" + data["CommentID"];
                let url = "/#Comment" + data["CommentID"];

                commentCard.innerHTML = `
                    <div class="commentAuthorImage">
                        <img src="https://api.multiavatar.com/${data["UserName"]}.svg" alt="">
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
                handleLikes()
                ExpandComments()
            }
        }
    });
}
CommentInputEventListenner()
ExpandComments()