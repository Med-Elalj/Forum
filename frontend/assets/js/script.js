// class testing 
const testing = document.querySelectorAll(".testing")
testing.forEach(elem => {
    elem.addEventListener('click', () => {
        console.log(elem.textContent);
    })
})
let popup = NaN;

//////////////////// Check User is Logged yet or not //////////////////
function checkUserIsLogged() {
    //const test = localStorage.setItem("token", "This is a test token")
    const token = localStorage.getItem("token")
    //////////// Here we have to send Token to Server
    /////////// to check its valid or not 
    return token !== null && token !== undefined;
}

//////////////////// Popup function //////////////////
const popUp = () => {
    const popupContainer = document.getElementById("popupContainer")
    const popupHTML = `
        <div id="popup" class="popup">
            <div class="popup-content">
                <h1>Thanks for trying</h1>
                <p>Log in or sign up to add comments, likes, dislikes, and more.</p>
                <a href="login.html"><button>Log in</button></a>
                <a href="register.html"><button>Sign up</button></a>
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
