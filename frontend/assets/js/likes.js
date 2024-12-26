/*
 <!-- Post Start Content Card -->
                <div class="post-card">
                    <!-- User image -->
                    <div class="ProfileImage tweet-img"
                        style="background-image: url('https://ui-avatars.com/api/?name={{$post.UserName}}')">
                    </div>

                    <div class="post-details">
                        <div class="row-tweet">
                            <div class="post-header">
                                <!-- Post Title -->
                                <span class="tweeter-name post" id="{{$post.ID}}">
                                    {{$post.Title}}
                                    <!-- Post Author Name And Date -->
                                    <br><span class="tweeter-handle">@{{$post.UserName}}
                                        {{$post.CreatedAt}}.</span>
                                </span>
                            </div>
                            {{if eq $.Profile.UserName $post.UserName}}
                            <!-- Control Posts -->
                            <div class="dropdown">
                                <i class="material-symbols-outlined">more_horiz</i>
                                <div class="content">
                                    <ul>
                                        <li><span class="material-symbols-outlined">edit</span>Edit</li>
                                        <li><span class="material-symbols-outlined">delete</span>Delete</li>
                                    </ul>
                                </div>
                            </div>
                            {{end}}
                        </div>
                        <!-- Post Content -->
                        <div class="post-content">
                            <p>{{$post.Content}}</p>
                        </div>
                        <span class="see-more">See More</span>

                        <!-- Post Categories -->
                        <div class="Hashtag">
                            {{range $post.Categories}}
                            <a href=""><span>#{{.}}</span></a>
                            {{end}}
                        </div>

                        <div class="post-footer">
                            <div class="react">
                                <!-- Post Like Counter -->
                                <div class="counters like" id="{{$post.ID}}">
                                    <i class="material-symbols-outlined popup-icon">thumb_up</i>
                                    <span>{{$post.LikeCount}}</span>
                                </div>
                                <!-- Post Dislike Counter -->
                                <div class="counters dislike" id="{{$post.ID}}">
                                    <i class="material-symbols-outlined popup-icon">thumb_down</i>
                                    <span>{{$post.DislikeCount}}</span>
                                </div>
                            </div>
                            <div class="comment post" id="{{$post.ID}}">
                                <!-- Post Comments Counter -->
                                <i class="material-symbols-outlined showCmnts">comment</i>
                                <span>10</span>
                            </div>
                        </div>
                    </div>
                </div>
                <!-- End Of Post Content Card -->

*/

// Handling Likes and Dislikes in both the frontend and backend

function handleLikes() {
    const likeBtns = document.querySelectorAll('.like');
    const dislikeBtns = document.querySelectorAll('.dislike');

    likeBtns.forEach(btn => {
        btn.addEventListener('click', async () => {
            const postId = btn.id;
            try {
                const res = await fetch(`/PostReaction`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ 
                        postId: postId,
                        "type": "like"
                     })
                });
                console.log(res.body);
                console.log(res);

                const data = await res.json();
                console.log(data);
                console.log("=====> Added = ", data.added);
                
               
                const dislike = btn.nextElementSibling
                if (data.added){
                    btn.classList.add("FILL")
                    dislike.classList.remove("FILL")
                }else{
                    btn.classList.remove("FILL")
                }
                dislike.querySelector('span').innerText = data.dislikes;
                btn.querySelector('span').innerText = data.likes;
            } catch (error) {
                console.log(error)
            }
          
        });
    });

    dislikeBtns.forEach(btn => {
        btn.addEventListener('click', async () => {
            const postId = btn.id;
            try{
                const res = await fetch(`/PostReaction`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ 
                        postId: postId,
                        "type": "dislike"
                     })
                });
    
                const data = await res.json();
                console.log(data);
                const like = btn.previousElementSibling
                console.log("=====> Added = ", data.added);
                
                if (data.added){
                    btn.classList.add("FILL")
                    like.classList.remove("FILL")
                }else{
                    btn.classList.remove("FILL")
                }
                like.querySelector('span').innerText = data.likes;
                btn.querySelector('span').innerText = data.dislikes;
            
            }catch(error){
                console.log(error)
            }   
         });
    });
}

handleLikes();
