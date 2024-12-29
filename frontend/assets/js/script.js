let popup = NaN;
const urlParams = new URLSearchParams(window.location.search);
const type = urlParams.get('type');

async function fetchPosts(offset, type) {
    type = type ? type : "home"
    let category_name = urlParams.get('category')
    console.log('Fetching posts with offset:', offset, type);
    const x = await fetch(`/infinite-scroll?offset=${offset}&type=${type}${category_name ? `&category=${category_name}` : ''}`)
        .then(response => response.json())
        .then(posts => {
            console.log('Fetched posts:', posts);
            const postsContainer = document.querySelector('.main-feed');
            posts.forEach(post => {
                const postCard = document.createElement('div');
                postCard.classList.add('post-card');
                console.log('Creating post card for post:', post);
                const profileImage = document.createElement('div');
                profileImage.className = 'ProfileImage tweet-img';
                profileImage.style.backgroundImage = `url('https://api.multiavatar.com/${post.author_username}.svg')`;

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
                postCard.append(profileImage, postDetails);
                console.log(' posts container:', postsContainer);
                console.log('posts container:', postCard);
                postsContainer.append(postCard);
            });
        });
        readPost()
        showLeftSidebarMobile()
        seeMore()
        handleLikes()
}

function infiniteScroll() {
    let offset = 10;
    let limit = 10;
    let timeout = null;
    window.addEventListener('scroll', () => {
        clearTimeout(timeout);
        timeout = setTimeout(() => {
            const { scrollTop, scrollHeight, clientHeight } = document.documentElement;
            console.log('Scroll event detected:', { scrollTop, scrollHeight, clientHeight });
            if (scrollTop + clientHeight >= scrollHeight - 5) {
                console.log('Fetching more posts...');
                fetchPosts(offset, type)
                    .then(() => {
                        offset += limit;
                        console.log('Finished fetching posts, new offset:', offset);
                    });
            }
        }, 1000);
    });

}
// get Query values from url to fetch the posts


infiniteScroll()
fetchPosts(0, type)


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

function showLeftSidebarMobile() {
    const sidebardLeft = document.querySelector('.sidebar-left')
    const menuIcon = document.querySelector('.menu')
    menuIcon.addEventListener('click', () => {
        sidebardLeft.style.left = sidebardLeft.style.left === '0px' ? '-100%' : '0px'
    })

    // Hiding Humberger Menu on Mobile if the user change the view port
    const x = window.matchMedia("(min-width: 768px)")
    x.addEventListener('change', (e) => {
        if (e.matches) {
            sidebardLeft.style.left = '2.5%'
        } else {
            sidebardLeft.style.left = '-100%'
        }
    })
}
//////////////////// See more Option  ////////////
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
        const html = await response.text();
        return html
    } catch (error) {
        console.error('Error fetching HTML:', error);
    }
}
function readPost() {
    let postButton = document.querySelectorAll(".post")
    postButton.forEach(elem => {
        elem.addEventListener('click', async () => {
            // Get id to send request to get Post : elem.id
            const html = await fetchPost(`/post/${elem.id}`)
            const postContent = document.querySelector('.postContainer')
            postContent.classList.remove("closed")
            if (!document.getElementById("ScriptInjected")) {
                const script = document.createElement("script")
                script.id = "ScriptInjected"
                script.src = "/assets/js/script-post.js"
                document.body.appendChild(script)
            }
            postContent.innerHTML = html
            // stop scrolling on background if the pop up opened
            document.body.classList.add("stop-scrolling");

            document.addEventListener('click', (event) => {
                if (event.target == postContent || event.target.classList.contains("close-post")) {
                    postContent.classList.add("closed")
                    //restore the scrolling on the background page :D 
                    document.body.classList.remove("stop-scrolling");
                    document.getElementById("ScriptInjected").remove()
                }
            })
            // recall Like.js to listen on Elemnts in post page
            handleLikes()
        })
    })

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

postControlList()
readPost()
showLeftSidebarMobile()
seeMore()

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
    
    // Observe DOM changes
    observer.observe(document.body, {
        childList: true,
        subtree: true
    });
    setInterval(updateAllTimes, 60000);
});
