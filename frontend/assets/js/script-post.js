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
                                <span isPost="false" class="like" id="${data["CommentID"]}">
                                    <i class="material-symbols-outlined">
                                        thumb_up
                                    </i>
                                    <span>0</span>
                                </span>
                                <span isPost="false" class="dislike" id="${data["CommentID"]}">
                                    <i class="material-symbols-outlined">thumb_down</i>
                                    <span>0</span>
                                </span>
                            </div>
                        </div>
                        <div class="commentInfo">
                            <p class="commentData collapse"></p>
                        </div>
                    </div>
                `;
                commentCard.querySelector('.commentData').innerText = data["Content"];
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
