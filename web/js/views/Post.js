import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("MyForum");
    }

    async getHtml() {
        return `
        <header id="top">
            <div class="row">
                <div class="column lpad">
                    <div class="logo">
                        <a href="http://localhost:8081/" data-link>MyForum</a>
                    </div>
                </div>
                <div class="column ar lpad">
                    <nav class="menu">
                        <a href="/" id="h_posts" class="current nav_link" data-link>Posts</a>
                    </nav>
                </div>
            </div>
        </header>
        <div class="row">
        </div>
        `;
    }

    async Init() {
        let id = window.location.pathname.replace("/post/", "");
        let url = new URL(`http://localhost:8080/post/${id}`);
        console.log("post worked", url.href)
        let app = document.getElementById('app');
        let container = document.createElement('div');
        container.setAttribute('id', 'container');
        app.appendChild(container);
        drawPostPage();
    }
}

async function drawPostPage() {
    let container = document.getElementById('container');
    // post contaner
    let post_container = document.createElement('div');
    post_container.setAttribute('class', 'post-container');
    let like_box = document.createElement('div');
    like_box.setAttribute('class', 'like-box');
    let like = document.createElement('div');
    like.setAttribute('id', 'like');
    like.innerHTML = 'like';
    let like_plus = document.createElement('button');
    like_plus.setAttribute('class', 'like-plus');
    let like_minus = document.createElement('button');
    like_minus.setAttribute('class', 'like-minus');
    let post_box = document.createElement('div');
    post_box.setAttribute('class', 'post-box');
    let title = document.createElement('div');
    title.setAttribute('id', 'title');
    title.innerHTML = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor.";
    let hr = document.createElement('hr');
    hr.setAttribute('class', 'hr-box');
    let post_body = document.createElement('div');
    post_body.setAttribute('id', 'post-body');
    post_body.innerHTML = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.";
    let author = document.createElement('div');
    author.setAttribute('id', 'author');
    author.innerHTML = 'Creted by Dos at 02.02.2021';
    like_box.appendChild(like_plus);
    like_box.appendChild(like);
    like_box.appendChild(like_minus);
    post_box.appendChild(title);
    post_box.appendChild(hr);
    post_box.appendChild(post_body);
    post_box.appendChild(author);
    post_container.appendChild(like_box);
    post_container.appendChild(post_box);
    //create-comment-contaner
    let create_comment_container = document.createElement('form');
    create_comment_container.setAttribute('class', 'create-comment-container');
    let post_id = document.createElement('input');
    post_id.setAttribute('type', 'hidden');
    post_id.setAttribute('id', "post-id");
    post_id.setAttribute('value', "");
    let create_comment_header = document.createElement('h2');
    create_comment_header.setAttribute('class', 'create-comment-header');
    create_comment_header.innerHTML = "Your Comment:";
    let textarea_form = document.createElement('textarea');
    textarea_form.setAttribute('id', 'textarea-form');
    textarea_form.setAttribute('autocapitalize', 'sentences');
    create_comment_container.appendChild(post_id);
    create_comment_container.appendChild(create_comment_header);
    create_comment_container.appendChild(textarea_form);
    let submit_comment = document.createElement('button');
    submit_comment.setAttribute('type', 'submit');
    submit_comment.setAttribute('class', 'submt-comment');
    submit_comment.innerHTML = "Add Your Comment";
    create_comment_container.appendChild(submit_comment);
    //comments container
    let comments_container = document.createElement('div');
    comments_container.setAttribute('class', 'comments-container');
    let all_comments = document.createElement('div');
    all_comments.setAttribute('class', 'all-comments');
    all_comments.innerHTML = "All Comments";
    comments_container.appendChild(all_comments);
    let hr_comments = document.createElement('hr');
    hr_comments.setAttribute('class', 'hr-box');
    comments_container.appendChild(hr_comments);
    let like_box_comment = document.createElement('div');
    like_box_comment.setAttribute('class', 'like-box-comment')
    let like_comment = document.createElement('div');
    like_comment.setAttribute('id', 'like-comment');
    like_comment.innerHTML = 'like';
    let like_plus_comment = document.createElement('button');
    like_plus_comment.setAttribute('class', 'like-plus-comment');
    let like_minus_comment = document.createElement('button');
    like_minus_comment.setAttribute('class', 'like-minus-comment');
    like_box_comment.appendChild(like_plus_comment);
    like_box_comment.appendChild(like_comment);
    like_box_comment.appendChild(like_minus_comment);
    comments_container.appendChild(like_box_comment);
    let comment_body = document.createElement('div');
    comment_body.setAttribute('class', 'comment-body');
    comment_body.innerHTML = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.";
    comments_container.appendChild(comment_body);
    let author_comment = document.createElement('div');
    author_comment.setAttribute('class', 'author-comment');
    author_comment.innerHTML = "Commented 8 aug 2021 at 8:20\nby Nurs";
    comments_container.appendChild(author_comment);
    //apend to post page container 
    container.appendChild(post_container);
    container.appendChild(create_comment_container);
    container.appendChild(comments_container);
}