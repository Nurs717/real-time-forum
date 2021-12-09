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
        drawPost();
    }
}

async function drawPost() {
    let container = document.getElementById('container');
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
    post_body.innerHTML = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.";
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
    container.appendChild(like_box);
    container.appendChild(post_box);
}