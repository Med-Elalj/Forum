function createPost() {
    // const CreatePostArea = document.querySelector(".new-post-header")
    const CreatePostModel = document.querySelector(".postModal")
    const closeCreatePostModal = document.querySelector(".titleInput .close-post")
    const ErrorBox = document.querySelector(".ErrorMessage")
    window.onclick = function (event) {
        if (event.target == CreatePostModel) {
            CreatePostModel.style.display = "none"
        } else if (event.target == popup) {
            popup.style.display = "none"
        }
    }

    const CreatePostInputTitle = document.querySelector(".titleInput input")
    CreatePostModel.style.display = "flex"
    CreatePostInputTitle.focus()
    closeCreatePostModal.addEventListener('click', () => {
        CreatePostModel.style.display = "none"
        document.getElementById("CreatePostScriptInjected").remove()
    })

    const form = document.querySelector('.CreatePostContainer form');
    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        const title = form.title.value;
        const content = form.content.value;
        const categories = Array.from(form.category).filter((input) => input.checked).map((input) => input.id);
        console.log(categories.length);
        
        if (categories.length === 0) {
            ErrorBox.style.display = "flex"
            document.querySelector(".message").innerText = "Please Select At Least One Category"
            setTimeout(function(){
                ErrorBox.style.display = "none"
            }, 5000)
            return;
        }
        const data = {
            title,
            content,
            categories
        };
        
        const response = await fetch('/createPost', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        });
        console.log("====>",response.status );
        
        if (response.status === 200) {
            
            const post = await response.json();
            const postCard = document.createElement('div');
            postCard.classList.add('post-card');
            postCard.innerHTML = `
                <div class="ProfileImage tweet-img" style="background-image: url('https://api.multiavatar.com/${post.UserName}.svg')"></div>
                <div class="post-details">
                    <div class="row-tweet">
                        <div class="post-header">
                            <span class="tweeter-name post" id="${post.ID}">
                                <span class="text-title"></span>
                                <br><span class="tweeter-handle">@${post.UserName} ${post.CreatedAt}.</span>
                            </span>
                        </div>
                    </div>
                    <div class="post-content">
                        <p></p>
                    </div>
                    <span class="see-more">See More</span>
                    <div class="Hashtag">
                        ${post.Categories.map((category) => `<a href=""><span>#${category}</span></a>`).join('')}
                    </div>
                    <div class="post-footer">
                        <div class="react">
                            <div class="counters like" id="${post.ID}">
                                <i class="material-symbols-outlined popup-icon">thumb_up</i>
                                <span>${post.LikeCount}</span>
                            </div>
                            <div class="counters dislike" id="${post.ID}">
                                <i class="material-symbols-outlined popup-icon">thumb_down</i>
                                <span>${post.DislikeCount}</span>
                            </div>
                        </div>
                        <div class="comment post" id="${post.ID}">
                            <i class="material-symbols-outlined showCmnts">comment</i>
                            <span>0</span>
                        </div>
                    </div>
                </div>
            `;
            // insert PostCard inside the main-feed and exact after first child of main-feed = new-tweet
            postCard.querySelector('.post-content p').innerText = post.Content
            postCard.querySelector('.text-title').innerText = post.Title
            const mainFeed = document.querySelector('.main-feed');
            postCard.classList.add('PostAdded');
            mainFeed.insertBefore(postCard, mainFeed.children[1]);
            CreatePostModel.style.display = "none";
            form.reset();
            // Recall Function To append new post to their Lestining Buttons
            seeMore()
            readPost() 
            handleLikes()
            createPostListner()
        }else{
            ErrorBox.style.display = "flex"
            document.querySelector(".message").innerText = "post with the same title and content already exist"
            setTimeout(function(){
                ErrorBox.style.display = "none"
            }, 5000)
            return;
        }
    });
}
