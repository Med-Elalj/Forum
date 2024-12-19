const popupContainer = document.getElementById('popupContainer');

function createPopup(html) {
    popupContainer.innerHTML = html;
    const popup = popupContainer.querySelector(':first-child');
    popup.style.display = 'flex';
}

function setupPopupTriggers(selector, html) {
    document.querySelectorAll(selector).forEach(trigger => {
        trigger.addEventListener('click', () => createPopup(html));
    });
}

const popups = [
    {
        selector: '.popup-icon',
        html: `
            <div id="popup" class="popup">
                <div class="popup-content">
                    <h1>Thanks for trying</h1>
                    <p>Log in or sign up to add comments, likes, dislikes, and more.</p>
                    <a href="login.html"><button>Log in</button></a>
                    <a href="register.html"><button>Sign up</button></a>
                    <a href="." class="logged-out">Stay logged out</a>
                </div>
            </div>`
    },
    {
        selector: '.textarea',
        html: `
            <div class="postModal" id="creatPost">
                <div class="createPostContainer">
                    <form action="/create-post" method="POST">
                        <input type="text" name="title" placeholder="Enter a Title" required maxlength="100">
                        <textarea name="content" placeholder="Share your thoughts..." required maxlength="1000"></textarea>
                        <div class="categories">
                            <input type="checkbox" name="category" id="general" value="General">
                            <label for="general">General</label>
                            <!-- Other category checkboxes -->
                        </div>
                        <input type="submit" value="Publish Post" class="submitButton">
                    </form>
                </div>
            </div>`
    },
    {
        selector: '.showCmnts',
        html: `
            <div class="postContainer">
                <!-- Simplified comments section -->
                <div class="postInfo">
                    <span class="Title">Sample Post</span>
                    <div class="postComments">
                        <!-- Minimal comment structure -->
                    </div>
                </div>
            </div>`
    }
];

// Setup popup triggers
popups.forEach(popup => setupPopupTriggers(popup.selector, popup.html));

// Single window click handler for closing popups
window.addEventListener('click', (event) => {
    const openPopup = document.querySelector('[style*="display: flex"]');
    if (openPopup && event.target === openPopup) {
        openPopup.style.display = 'none';
    }
});