let popup = NaN;

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

function readPost() {
    async function fetchPost(url) {
        try {
            //?id=${idValueFromClikedArea}
            const response = await fetch(url);
            const html = await response.text();
            return html
        } catch (error) {
            console.error('Error fetching HTML:', error);
        }
    }

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
const themeToggle = document.getElementById('switch');
const body = document.body;

function toggleDarkMode(isDark) {
    body.classList.toggle('dark-mode', isDark);
    localStorage.setItem('darkMode', isDark);
    body.classList.add('theme-transitioning');
    body.classList.remove('theme-transitioning')
}
const darkModeStored = localStorage.getItem('darkMode') === 'true';
toggleDarkMode(darkModeStored);
console.log("test", darkModeStored);

themeToggle.checked = darkModeStored;
themeToggle.addEventListener('change', (e) => toggleDarkMode(e.target.checked));



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


/* 

func InfiniteScroll(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil && err.Error() != "http: named cookie not present" {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized"+err.Error()))
		fmt.Println(err)
		return
	}
	if err != nil {
		c = &http.Cookie{}
	}

	uid, err := database.GetUidFromToken(DB, c.Value)
	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	offset_str := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(offset_str)
	if err != nil {
		offset = 0
	}
	fmt.Println("Offset:", offset)
	posts, err := database.QuerryLatestPosts(DB, uid, structs.Limit, offset)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, err)
		return
	}
	// Set the content type header to application/json
	w.Header().Add("Content-Type", "application/json")

	// Optionally set the status code to 200 OK
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(posts)
	fmt.Println(err)
}


*/
// Create Function to fetch new post in the user scrolling at the end of the page
// using the offset and limit to fetch the new post
// check function above from backend

async function fetchPosts(offset) {
    console.log('Fetching posts with offset:', offset);
    const x = await fetch(`/infinite-scroll?offset=${offset}`)
        .then(response => response.json())
        .then(posts => {
            console.log('Fetched posts:', posts);
            const postsContainer = document.querySelector('.main-feed');
            posts.forEach(post => {
                const postCard = document.createElement('div');
                postCard.classList.add('post-card');
                console.log('Creating post card for post:', post);
                postCard.innerHTML = `
                    <div class="ProfileImage tweet-img" style="background-image: url('https://api.multiavatar.com/${post.author_username}.svg')">
                    </div>
                    <div class="post-details">
                        <div class="row-tweet">
                            <div class="post-header">
                                <span class="tweeter-name post" id="${post.post_id}">
                                    ${post.post_title}
                                    <br><span class="tweeter-handle">@${post.author_username} ${post.post_creation_time}.</span>
                                </span>
                            </div>
                            <div class="dropdown">
                                <i class="material-symbols-outlined">more_horiz</i>
                                <div class="content">
                                    <ul>
                                        <li><span class="material-symbols-outlined">edit</span>Edit</li>
                                        <li><span class="material-symbols-outlined">delete</span>Delete</li>
                                    </ul>
                                </div>
                            </div>
                        </div>
                        <div class="post-content">
                            <p>${post.post_content}</p>
                        </div>
                        <span class="see-more">See More</span>
                        <div class="Hashtag">
                            ${post.post_categories ? post.post_categories.map(category => `<a href=""><span>#${category}</span></a>`).join('') : ''}
                        </div>
                        <div class="post-footer">
                            <div class="react" id="${post.ID}">
                                <div isPost="true" class='counters like  ${post.view && post.view === '1' ? 'FILL' : ''}' id="${post.post_id}">
                                    <i class="material-symbols-outlined popup-icon " id="${post.ID}">thumb_up</i>
                                    <span id="${post.post_id}">${post.like_count}</span>
                                </div>
                                <div isPost="true" class='counters dislike ${post.view && post.view === '0' ? 'FILL' : ''}' id="${post.post_id}">
                                    <i class="material-symbols-outlined popup-icon" id="${post.ID}">thumb_down</i>
                                    <span id="${post.post_id}">${post.dislike_count}</span>
                                </div>
                            </div>
                            <div class="comment post" id="${post.post_id}">
                                <i class="material-symbols-outlined showCmnts">comment</i>
                                <span>${post.comment_count}</span>
                            </div>
                        </div>
                    </div>
                `;
                console.log(' posts container:', postsContainer);
                console.log('posts container:', postCard);
                postsContainer.append(postCard);
            });
        });
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
                fetchPosts(offset)
                    .then(() => {
                        offset += limit;
                        console.log('Finished fetching posts, new offset:', offset);
                    });
            }
        }, 1000);
    });
}

infiniteScroll();
