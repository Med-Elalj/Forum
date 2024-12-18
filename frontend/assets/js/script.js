//###//////////////////// Start Listning dropDown Buttons ////////////
const CreatePostModel = document.querySelector(".postModal")
const CreatePostArea = document.querySelector(".new-tweet-details")
let popup = NaN;
const CreatePostInputTitle = document.querySelector(".titleInput input")
const dropdown = document.querySelectorAll('.dropdown i')
let contentList = document.querySelectorAll('.content')

dropdown.forEach(drop => {
    
    let contentSibling = drop.nextElementSibling
    drop.addEventListener('click', ()=>{
        contentSibling.classList.toggle("show")
    })
    document.addEventListener('click', function(event) {
        if (!contentSibling.contains(event.target) && !drop.contains(event.target) && contentSibling.classList.contains("show")) {
            console.log(contentSibling.classList);
            contentSibling.classList.remove('show');
        }
    });
})

//###//////////////////// END Listning dropDown Buttons ////////////


//###//////////////////// Start Listning on Like Dislike Buttons ////////////
let like = document.querySelectorAll(".react .like")
let dislike = document.querySelectorAll(".react .dislike")

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


const popUp = () => {
    const popupContainer = document.getElementById("popupContainer")
    popupContainer.innerHTML = popupHTML;
    popup = document.getElementById('popup');

    popup.style.display = 'flex';
}

window.onclick = function (event) {

    if (event.target == CreatePostModel){
        CreatePostModel.style.display = "none"
    } else if (event.target == popup){
        popup.style.display = "none"
    }
}
// Handling Like Button Clicked in Post Comments
like.forEach(like_elem => {
    like_elem.addEventListener('click', function(){
        // Check if the user is loggedin or not :
        if (checkUserIsLogged()){
            // Check from DB if like or dislike exits
            /// /// / // / / / / / /
            let dislike_sibling = like_elem.nextElementSibling;
            like_elem.classList.toggle("FILL");
            dislike_sibling.classList.remove("FILL");
        }else{
            popUp();
        }
    })
})

dislike.forEach(dislike_elem => {
    dislike_elem.addEventListener('click', function(){
        if (checkUserIsLogged()){
            let like_sibling = dislike_elem.previousElementSibling;
            dislike_elem.classList.toggle("FILL");
            like_sibling.classList.remove("FILL");
        }else{
            popUp();
        }
      

        // the  rest of the code will write here to send request
        // to backend to update database 
        
    })
})

//###//////////////////// Check User is Logged yet or not //////////////////
function checkUserIsLogged(){
    //const test = localStorage.setItem("token", "This is a test token")
    const token = localStorage.getItem("token")
    //////////// Here we have to send Token to Server
    /////////// to check its valid or not 
    return token !== null && token !== undefined;
}
//###//////////////////// END OF Listning on Like Dislike Buttons ////////////




//###//////////////////// See more Option  ////////////
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

//###////////////////////   Listening on user request of post  /////
// ///To Read Post We need to make request to backend, to get Full Page to Display it to the user

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
    elem.addEventListener('click', async ()=>{
        // Get id to send request to get Post : elem.id
        const html = await fetchPost(`/post?id=${elem.id}`)
        const postContent = document.querySelector('.postContainer')
        postContent.classList.remove("closed")
        if (!document.getElementById("ScriptInjected")){
            const script = document.createElement("script")
            script.id = "ScriptInjected"
            script.src  = "/assets/js/script-post.js"
            document.body.appendChild(script)
        }
        postContent.innerHTML = html
        // stop scrolling on background if the pop up opened
        document.body.classList.add("stop-scrolling");

        document.addEventListener('click', (event) =>{
            if (event.target == postContent || event.target.classList.contains("close-post")){
                postContent.classList.add("closed")
                //restore the scrolling on the background page :D 
                document.body.classList.remove("stop-scrolling");
            }
        })
    })
})



//###////////////////////  Creat NEW POST Listenner  //////////


CreatePostArea.addEventListener('click', ()=>{
    CreatePostModel.style.display  = "flex"
    CreatePostInputTitle.focus()
})