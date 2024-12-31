const commentInput = document.querySelector('.commentInput input');
const send_comment = document.querySelector('.send-comment');
const commentButton = document.querySelector('.CommentButton');
const postButton = document.querySelector('.PostButton');
const commentSection = document.querySelector('.postComments');
const postSection = document.querySelector('.ProfileAndPost');

function removeExpandCommentListeners() {
    let comments = document.querySelectorAll(".commentData")
    comments.forEach(elem => {
        elem.removeEventListener('click', ExpandComments);
    })
}

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

async function handleCommentEvent(e) {

    if (e.type === 'click' || (e.type === 'keypress' && e.key === 'Enter')) {
        e.preventDefault();
        const commentInput = e.target.closest('.commentInput').querySelector('input');
        const comment = commentInput.value;
        if (comment.trim() === '' || comment.length == 0)
            return;

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
        if (response.status == 401) {
            popUp();
            return;
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
            document.querySelector('.commentCount').textContent = data.CommentCount
            window.location.replace(url);
            

            // remove old Listeners
            removeHandeLikeListeners();
            // call new listeners
            handleLikes();

            // remove old Listners :
            removeExpandCommentListeners()
            // call new Listners
            ExpandComments();
        }
    }
}

function addCommentOnPressKey(e) {
    if (!throttleTimeout && e.key == "Enter") {
        handleCommentEvent(e);
        throttleTimeout = true
    }
    setTimeout(() => {
        throttleTimeout = false
    }, 1000);
}
let throttleTimeout = false
function addCommentOnClick(e){
    if (!throttleTimeout) {
        handleCommentEvent(e);
        throttleTimeout = true
    }
    setTimeout(() => {
        throttleTimeout = false
    }, 1000);
}

function removeCommentListtner()
{
    commentInput.removeEventListener('keypress', addCommentOnPressKey);
    send_comment.removeEventListener('click',  addCommentOnClick);
}
function CommentInputEventListenner() {
    commentInput.addEventListener('keypress', addCommentOnPressKey);
    send_comment.addEventListener('click',  addCommentOnClick);
}

function DisplayPost(){
    commentSection.style.display = 'flex';
    postSection.style.display = 'none';
}

function DisplayComments(){
    commentSection.style.display = 'none';
    postSection.style.display = 'flex';
}

function removePostButtonSwitcherListners(){
    commentButton.removeEventListener('click', DisplayPost);
    postButton.removeEventListener('click', DisplayComments);
}

function PostButtonSwitcher(){
    commentButton.addEventListener('click', DisplayPost);
    postButton.addEventListener('click', DisplayComments);
}
PostButtonSwitcher()
CommentInputEventListenner()
ExpandComments()


