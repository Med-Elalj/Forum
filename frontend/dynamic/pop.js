const popupContainer = document.getElementById('popupContainer');

const popupHTML = `
    <div id="popup" class="popup">
        <div class="popup-content">
            <h1>Thanks for trying</h1>
            <p>Log in or sign up to add comments, likes, dislikes, and more.</p>
            <a href="login.html"><button>Log in</button></a>
            <a href="register.html"><button>Sign up</button></a>
            <a href="." class="logged-out">Stay logged out</a>
        </div>
    </div>
`;

const popupIcons = document.querySelectorAll('.popup-icon');

popupIcons.forEach(function (popupIcon) {
    popupIcon.addEventListener('click', function () {
        popupContainer.innerHTML = popupHTML;

        const popup = document.getElementById('popup');
        popup.style.display = 'flex';
    });
});

window.onclick = function (event) {
    const popup = document.getElementById('popup');
    if (event.target === popup) {
        popup.style.display = 'none';
    }
}
