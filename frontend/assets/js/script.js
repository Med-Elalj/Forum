let popup = NaN;
const UrlParams = new URLSearchParams(window.location.search);
const sidebardLeft = document.querySelector('.sidebar-left')
const menuIcon = document.querySelector('.menu')
const windowMedia = window.matchMedia("(min-width: 768px)")
const type = UrlParams.get('type');
const username = UrlParams.get('username');
let errorr = ""
async function fetchPosts(offset, type) {
    type = type ? type : "home"
    let category_name = UrlParams.get('category')
    const postsContainer = document.querySelector('.main-feed');
    const x = await fetch(`/infinite-scroll?offset=${offset}&type=${type}${category_name ? `&category=${category_name}` : ''}${username ? `&username=${username}`:''}`)
        .then(response => response.json())
        .then(posts => {
            if (type === 'profile') {
                console.log(posts.profile);
                const pImage = document.querySelector('.profileImage img')
                const pName = document.querySelector('.profileName')
                const pCounts = document.querySelector('.posts .postCounts')
                const cCounts = document.querySelector('.comments .postCounts')
                pImage.src = `https://api.multiavatar.com/${posts.profile.UserName}.svg`
                pName.textContent = posts.profile.UserName
                pCounts.textContent = `${posts.profile.ArticleCount} Articles`
                cCounts.textContent = `${posts.profile.CommentCount} Comments`
            }
            posts.posts.forEach(post => {
                const postCard = document.createElement('div');
                postCard.classList.add('post-card');
                const profileLink =  document.createElement('a')
                profileLink.href = `/?type=profile&username=${post.author_username}`
                const profileImage = document.createElement('div');
                profileImage.className = 'ProfileImage tweet-img';
                profileImage.style.backgroundImage = `url('https://api.multiavatar.com/${post.author_username}.svg')`;
                profileLink.appendChild(profileImage)

                const postDetails = document.createElement('div');
                postDetails.className = 'post-details';

                const rowTweet = document.createElement('div');
                rowTweet.className = 'row-tweet';

                const postHeader = document.createElement('div');
                postHeader.className = 'post-header';

                const tweeterName = document.createElement('span');
                tweeterName.className = 'tweeter-name post';
                tweeterName.id = post.post_id;
                tweeterName.innerHTML = `${post.post_title.replace(/</g, "&lt;").replace(/>/g, "&gt;")}<br>
                    <span class="tweeter-handle">@${post.author_username}</span>
                    <span class="material-symbols-outlined" id="timer">schedule</span>
                    <span class="post-time" data-time="${post.post_creation_time}"> ${post.post_creation_time}</span>
                    `;

               
                // const dropdown = document.createElement('div');
                // dropdown.className = 'dropdown';

                // const dropdownIcon = document.createElement('i');
                // dropdownIcon.className = 'material-symbols-outlined';
                // dropdownIcon.textContent = 'more_horiz';

                // const dropdownContent = document.createElement('div');
                // dropdownContent.className = 'content';

                // const dropdownList = document.createElement('ul');

                // const editItem = document.createElement('li');
                // editItem.innerHTML = '<span class="material-symbols-outlined">edit</span>Edit';

                // const deleteItem = document.createElement('li');
                // deleteItem.innerHTML = '<span class="material-symbols-outlined">delete</span>Delete';

                // dropdownList.append(editItem, deleteItem);
                // dropdownContent.appendChild(dropdownList);
                // dropdown.append(dropdownIcon, dropdownContent);

                postHeader.appendChild(tweeterName);
                rowTweet.append(postHeader);

                const postContent = document.createElement('div');
                const postParagraph = document.createElement('p');
                postContent.className = 'post-content';
                postParagraph.innerText = `${post.post_content}`
                postContent.appendChild(postParagraph);

                const seeMore = document.createElement('span');
                seeMore.className = 'see-more';
                seeMore.textContent = 'See More';

                const hashtag = document.createElement('div');
                hashtag.className = 'Hashtag';
                if (post.post_categories) {
                    post.post_categories.forEach(category => {
                        const categoryLink = document.createElement('a');
                        categoryLink.href = '/?type=category&&category=' + category;
                        categoryLink.innerHTML = `<span>#${category}</span>`;
                        hashtag.appendChild(categoryLink);
                    });
                }

                const postFooter = document.createElement('div');
                postFooter.className = 'post-footer';

                const react = document.createElement('div');
                react.className = 'react';
                react.id = post.ID;

                const likeCounter = document.createElement('div');
                likeCounter.setAttribute('isPost', 'true');
                likeCounter.className = `counters like ${post.view && post.view === '1' ? 'FILL' : ''}`;
                likeCounter.id = post.post_id;
                likeCounter.innerHTML = `<i class="material-symbols-outlined popup-icon" id="${post.ID}">thumb_up</i><span id="${post.post_id}">${post.like_count}</span>`;

                const dislikeCounter = document.createElement('div');
                dislikeCounter.setAttribute('isPost', 'true');
                dislikeCounter.className = `counters dislike ${post.view && post.view === '0' ? 'FILL' : ''}`;
                dislikeCounter.id = post.post_id;
                dislikeCounter.innerHTML = `<i class="material-symbols-outlined popup-icon" id="${post.ID}">thumb_down</i><span id="${post.post_id}">${post.dislike_count}</span>`;

                react.append(likeCounter, dislikeCounter);

                const comment = document.createElement('div');
                comment.className = 'comment post';
                comment.id = post.post_id;
                comment.innerHTML = `<i class="material-symbols-outlined showCmnts">comment</i><span>${post.comment_count}</span>`;

                postFooter.append(react, comment);

                postDetails.append(rowTweet, postContent, seeMore, hashtag, postFooter);
                postCard.append(profileLink, postDetails);
                postsContainer.append(postCard);
            });
        }).catch(error => {
            console.log(error);
            
            //  window.location.href = '/error?code=404&message=Page Not Found';
        }
        );
        handleLikes(false)
        handleLikes(true)
        
        removeReadPostListner()
        readPost()

        removeShowLeftSidebarMobileListner()
        showLeftSidebarMobile()

        removeSeeMoreListner()
        seeMore()


}

function infiniteScroll() {
    let offset = 10;
    let limit = 10;
    let timeout = null;
    window.addEventListener('scroll', () => {
        clearTimeout(timeout);
        timeout = setTimeout(async () => {
            const { scrollTop, scrollHeight, clientHeight } = document.documentElement;
            if (scrollTop + clientHeight >= scrollHeight - 5) {
                await fetchPosts(offset, type)
            }
        }, 1000);
    });

}
// get Query values from url to fetch the posts



//////////////////// Popup function //////////////////
const popUp = () => {
    const popupContainer = document.getElementById("popupContainer")
    const popupHTML = `
        <div id="popup" class="popup">
            <div class="popup-content">
                <h1>Thanks for trying</h1>
                <p>Log in or sign up to add comments, likes, dislikes, and more.</p>
                <a href="/login"><button>Log in</button></a>
                <a href="/register"><button>Sign up</button></a>
                <span class="logged-out">Stay logged out</span>
            </div>
        </div>
    `;
    popupContainer.innerHTML = popupHTML;
    popup = document.getElementById('popup');

    popup.style.display = 'flex';
    popupContainer.addEventListener('click', (e) => {
        if (e.target === popup || e.target.classList.contains('logged-out')) {
            popup.style.display = 'none';
        }
    })
}

document.querySelectorAll('.nav-mobile a div').forEach(function (div) {
    div.addEventListener('click', function () {
        // Remove 'clicked' class from all divs
        document.querySelectorAll('.nav-mobile a div').forEach(function (item) {
            item.classList.remove('clicked');
        });

        // Add the 'clicked' class to the clicked div to show the underline
        this.classList.add('clicked');
    });
});

//////////////// Start Listning dropDown List For Posts ////////////
function postControlList() {

    const dropdown = document.querySelectorAll('.dropdown i, .dropdown .ProfileImage')
    
    dropdown.forEach(drop => {
        let contentSibling = drop.nextElementSibling
        drop.addEventListener('click', () => {
            contentSibling.classList.toggle("show")
        })
        document.addEventListener('click', function (event) {
            if (!contentSibling.contains(event.target) && !drop.contains(event.target) && contentSibling.classList.contains("show")) {
                contentSibling.classList.remove('show');
            }
        });
    })
}
//////////////////// Menu Icon On header PAGE Burger Icon for Mobile //////////

function showAndHideSideBar(e) {
    const commentSection = document.querySelector('.postComments')
    const postSection = document.querySelector('.ProfileAndPost')
    if (e.matches) {
        sidebardLeft.style.left = "2.5%"
        if (commentSection)
            commentSection.style.display = 'flex';
        if (postSection)
            postSection.style.display = 'flex';
    } else {
        if (commentSection)
            commentSection.style.display = 'none';
        if (postSection)
            postSection.style.display = 'flex';
        sidebardLeft.style.left = '-100%'
    }
   
}
function MenuIcon(){
    sidebardLeft.style.left = sidebardLeft.style.left === '0%' ? '-100%' : '0%'
}

function removeShowLeftSidebarMobileListner() {
    menuIcon.removeEventListener('click', MenuIcon)
    windowMedia.removeEventListener('change', showAndHideSideBar)
}

function showLeftSidebarMobile() {
    // Hiding Humberger Menu on Mobile if the user change the view port
    windowMedia.addEventListener('change', showAndHideSideBar)
    menuIcon.addEventListener('click', MenuIcon)
}
//////////////////// See more Option  ////////////
function removeSeeMoreListner() {
    document.querySelectorAll('.see-more').forEach(tweetText => {
        const seeMoreLink = tweetText;
        seeMoreLink.removeEventListener('click', seeMore);
    })
}

function seeMore() {
    document.querySelectorAll('.see-more').forEach(tweetText => {
        const seeMoreLink = tweetText;

        // Only show 'See More' if text is actually truncated
        const paragraph = tweetText.previousElementSibling.querySelector('p');
        if (paragraph.scrollHeight <= 50) {
            seeMoreLink.style.display = 'none';
        }

        seeMoreLink.addEventListener('click', () => {
            tweetText.previousElementSibling.classList.toggle('expanded');

            // Toggle see more/see less text
            seeMoreLink.textContent = tweetText.previousElementSibling.classList.contains('expanded')
                ? 'See Less'
                : 'See More';
        });
    });
}
////////////////////   Listening on user request of post  /////
// ///To Read Post We need to make request to backend, to get Full Page to Display it to the user

async function fetchPost(url) {
    try {
        const response = await fetch(url);
        if (response.status != 200)
        {
            window.location.href = '/error?code=404&message=Page Not Found';
            return false
        }
        const html = await response.text();
        return html
    } catch (error) {
        errorr = error
    }
}
function removeReadPostListner() {
    let postButton = document.querySelectorAll(".post")
    postButton.forEach(elem => {
        elem.removeEventListener('click',  loadPostContent(elem));
    })
}
function readPost() {
    let postButton = document.querySelectorAll(".post")

    postButton.forEach(elem => {
        elem.addEventListener('click', loadPostContent(elem))
    })
}


function loadPostContent(elem) {
    return async () => {
        // Get id to send request to get Post : elem.id
        const html = await fetchPost(`/post/${elem.id}`);
        if (!html) return;
        const postContent = document.querySelector('.postContainer');

        postContent.innerHTML = html;
        // stop scrolling on background if the pop up opened
        if (!document.getElementById("ScriptInjected")) {
            const script = document.createElement("script");
            script.id = "ScriptInjected";
            script.src = "/assets/js/comments.js";
            document.body.appendChild(script);
        }
        document.body.classList.add("stop-scrolling");

        document.addEventListener('click', (event) => {
            if (event.target == postContent || event.target.classList.contains("close-post")) {
                ExpandComments(false);
                CommentInputEventListenner(false);
                PostButtonSwitcher(false);
                postContent.innerHTML = "";
                postContent.classList.add("closed");
                //restore the scrolling on the background page :D 
                document.body.classList.remove("stop-scrolling");
                if (document.getElementById("ScriptInjected"))
                    document.getElementById("ScriptInjected").remove();
            
                console.log("Close Post");
            }
        });
        postContent.classList.remove("closed");

        // recall Like.js to listen on Elemnts in post page
        ListenOncommentButtom(false);
        ListenOncommentButtom(true);
        handleLikes(false);
        handleLikes(true);
    };
}

function DisplayPost(){
    const commentSection = document.querySelector('.postComments');
    const postSection = document.querySelector('.ProfileAndPost');
    if (!windowMedia.matches){
        commentSection.style.display = 'flex';
        postSection.style.display = 'none';
    }
}

function ListenOncommentButtom(add){
    const commentButton = document.querySelector('.CommentButton');
    if (add){
        commentButton.addEventListener('click', DisplayPost);
    }else{
        commentButton.removeEventListener('click', DisplayPost);

    }
}
// Dark Mode && save token in local storage
const themeToggle = document.querySelectorAll('#switch');
const body = document.body;

function toggleDarkMode(isDark) {
    body.classList.toggle('dark-mode', isDark);
    localStorage.setItem('darkMode', isDark);
    body.classList.add('theme-transitioning');
    body.classList.remove('theme-transitioning')
}
const darkModeStored = localStorage.getItem('darkMode') === 'true';
toggleDarkMode(darkModeStored);

themeToggle.forEach(elem => {
    elem.checked = darkModeStored;
    elem.addEventListener('change', () => {
        toggleDarkMode(elem.checked)
    });
})

document.querySelectorAll('.nav-mobile a div').forEach(function(div) {
    div.addEventListener('click', function() {
        // Remove 'clicked' class from all divs
        document.querySelectorAll('.nav-mobile a div').forEach(function(item) {
            item.classList.remove('clicked');
        });

        // Add the 'clicked' class to the clicked div to show the underline
        this.classList.add('clicked');
    });
});


// navbar for phone and ipad
window.addEventListener('load', () => {
    if (!window.location.hash) {
        window.location.hash = '#posts';
    }
});

window.addEventListener('hashchange', () => {
    const hash = window.location.hash;
    document.querySelectorAll('#posts, #categories').forEach(section => {
        section.style.display === '#posts' ? 'block' : 'none';
    });
});

// change the time to be more readable
function timeAgo(date) {
    const seconds = Math.floor((new Date() - new Date(date)) / 1000);
    
    if (seconds < 60) return 'just now';
    if (seconds < 3600) return `${Math.floor(seconds/60)}m ago`;
    if (seconds < 86400) return `${Math.floor(seconds/3600)}h ago`;
    if (seconds < 604800) return `${Math.floor(seconds/86400)}d ago`;
    if (seconds < 2592000) return `${Math.floor(seconds/604800)}w ago`;
    if (seconds < 31536000) return `${Math.floor(seconds/2592000)}mo ago`;
    return `${Math.floor(seconds/31536000)}y ago`;
}

function updateAllTimes() {
    const timeElements = document.querySelectorAll('.post-time, .commentTime, .postDate');
    timeElements.forEach(el => {
        if (el.dataset.time) {
            el.textContent = timeAgo(el.dataset.time);
        }
    });
}
const observer = new MutationObserver(() => {
    updateAllTimes();
});

document.addEventListener('DOMContentLoaded', () => {
    updateAllTimes();
    observer.observe(document.body, {
        childList: true,
        subtree: true
    });
    setInterval(updateAllTimes, 60000);
});


infiniteScroll()
fetchPosts(0, type)
postControlList()
readPost()
showLeftSidebarMobile()
seeMore()
