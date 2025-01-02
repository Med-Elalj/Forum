
function toggleCollapse(elem, comments) {
    console.log("Clic,ed");

    elem.classList.toggle("collapse");
    comments.forEach(second_elem => {
        if (second_elem != elem)
            second_elem.classList.add("collapse");
    });
}
function ExpandComments(flag) {
    // Expand Comment and read Content...
    let comments = document.querySelectorAll(".commentData")
    comments.forEach(elem => {
        if (flag) {
            elem.addEventListener('click', () => toggleCollapse(elem, comments))

        } else {
            elem.removeEventListener('click', () => toggleCollapse(elem, comments))
        }
    })
}

function CommentErrorMsg(msg) {
    const commentError = document.querySelector('.CommentErrorMessage');
    commentError.style.display = "block"
    commentError.innerText = msg;
    setTimeout(() => {
        commentError.style.display = "none"
        commentError.innerText = "";
    }, 5000);
}

// Remove duplicate 500 status check
async function handleCommentEvent(e) {

    if (e.type === 'click' || (e.type === 'keypress' && e.key === 'Enter')) {
        e.preventDefault();
        const commentValue = e.target.closest('.commentInput').querySelector('input');

        const comment = commentValue.value;
        if (comment.trim() === '' || comment.length === 0) {
            return;
        }

        const postID = commentValue.id;
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

        if (response.status === 401) {
            popUp();
            return;
        }
        if (response.status === 429) {
            CommentErrorMsg(`Slow down! Good comments take time—quality over speed! try again after 1 minute 😊`);
            return;
        }
        if (response.status === 500) {
            CommentErrorMsg("Oops! It looks like you've already posted this comment. Please try something new!");
            return;
        }
        if (response.status == 500) {
            commentError.style.display = "block"
            commentError.innerText = "Oops! It looks like you've already posted this comment. Please try something new!";
            setTimeout(() => {
                commentError.style.display = "none"
                commentError.innerText = "";
            }, 5000);
            return;
        }
        const data = await response.json();
        commentValue.value = '';
        if (data["status"] == "ok") {
            const commentContainer = document.querySelector('.Comments');
            const commentCard = document.createElement('div');
            commentCard.classList.add('commentCard');
            commentCard.classList.add('CommentAdded');

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
            document.querySelector('.commentCount').textContent = data.CommentCount
            // remove old Listeners
            handleLikes(false);
            // call new listeners
            handleLikes(true);

            // remove old Listners :
            ExpandComments(false)
            // call new Listners
            ExpandComments(true);
        }
    }
}

function CommentInputEventListenner(flag) {
    const send_comment = document.querySelector('.send-comment');
    const commentInput = document.querySelector('.commentInput input');
    if (flag) {
        commentInput.addEventListener('keypress', handleCommentEvent);
        send_comment.addEventListener('click', handleCommentEvent);
    } else {
        commentInput.removeEventListener('keypress', handleCommentEvent);
        send_comment.removeEventListener('click', handleCommentEvent);
    }
}

function DisplayComments() {

    const commentSection = document.querySelector('.postComments');
    const postSection = document.querySelector('.ProfileAndPost');
    commentSection.style.display = 'none';
    postSection.style.display = 'flex';
}

function PostButtonSwitcher(flag) {

    const postButton = document.querySelector('.PostButton');

    if (flag) {

        postButton.addEventListener('click', DisplayComments);
    } else {
        postButton.removeEventListener('click', DisplayComments);
    }
}
PostButtonSwitcher(true)
CommentInputEventListenner(true)
ExpandComments(true)