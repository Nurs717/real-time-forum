import AbstractView from "./AbstractView.js";

async function getPosts(url) {
    // const url = new URL(window.location.href);
    fetch(url.href, {
        method: "GET",
        credentials: 'include'
    }).then(
        (response) => {
            if (response.ok) {
                // draw nav menu
                drawNavMenuLoggedIn();
                //draw username
                drawWelcomUserName();
            } else {
                // draw nav menu logged out
                drawNavmenuLoggedOut();
            }
            //draw categories
            drawCategories();
            //draw posts
            drawPosts();
            return response.json()
        }
    ).then(
        (data) => {
            console.log('posts:', data);
            let app = document.getElementById("posts");
            if (document.cookie != "") {
                let username = document.getElementById('welcome_username');
                username.innerHTML = data[0].username;
            }
            data.map(post => {
                let line = document.createElement('p');
                let line2 = document.createElement('p');
                line.innerText = 'username: ' + post.username;
                line2.innerText = ' title: ' + post.post_title;
                app.appendChild(line);
                app.appendChild(line2)
            })
        }
    )
}

async function drawNavMenuLoggedIn() {
    let menu = document.getElementsByClassName("menu");
    let newPost = document.createElement('a');
    newPost.setAttribute('href', '/create-post');
    newPost.setAttribute('data-link', '');
    newPost.innerHTML = 'New Post';
    let logout = document.createElement('button');
    logout.setAttribute('logout-button', '');
    logout.innerHTML = 'Log Out';
    menu[0].appendChild(newPost);
    menu[0].appendChild(logout);
}

async function drawNavmenuLoggedOut() {
    let menu = document.getElementsByClassName("menu");
    let register = document.createElement('a');
    register.setAttribute('href', '/signup');
    register.setAttribute('data-link', '');
    register.innerHTML = 'Register';
    let login = document.createElement('a');
    login.setAttribute('href', '/login');
    login.setAttribute('data-link', '');
    login.innerHTML = 'Log In';
    menu[0].appendChild(register);
    menu[0].appendChild(login);
}

async function drawWelcomUserName() {
    let row = document.getElementsByClassName("row");
    let username = document.createElement('div');
    username.setAttribute('class', 'column lpad top-msg ar');
    username.innerHTML = 'Welcome, ';
    let welcome = document.createElement('a');
    welcome.setAttribute('id', 'welcome_username')
    welcome.setAttribute('class', 'underline');
    username.appendChild(welcome);
    row[1].appendChild(username);
}

async function drawCategories() {
    let row = document.getElementsByClassName('row')
    let categories = document.createElement('div');
    categories.setAttribute('class', 'column lpad top-msg breadcrumb');
    categories.setAttribute('id', 'breadcrumb');
    let sport = document.createElement('button');
    sport.setAttribute('value', 'sport');
    sport.setAttribute('category-button', '');
    sport.innerHTML = 'sport';
    let religion = document.createElement('button');
    religion.setAttribute('value', 'religion');
    religion.setAttribute('category-button', '');
    religion.innerHTML = 'religion';
    let programming = document.createElement('button');
    programming.setAttribute('category-button', '');
    programming.setAttribute('value', 'programming');
    programming.innerHTML = 'programming'
    categories.appendChild(sport);
    categories.appendChild(religion);
    categories.appendChild(programming);
    row[1].appendChild(categories);
}

async function drawPosts() {
    let app = document.getElementById("app");
    let posts = document.createElement("div");
    posts.setAttribute('id', 'posts');
    posts.setAttribute('class', 'row mt')
    app.appendChild(posts);
}

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("MyForum");
        this.myUrl = "http://localhost:8080/";
    }

    async getHtml() {
        return `
        <header id="top">
            <div class="row">
                <div class="column lpad">
                    <div class="logo">
                        <span>MyForum</span>
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
        let url = new URL("http://localhost:8080/");
        getPosts(url);

        const button = document.getElementById("app");
        button.addEventListener("click", async(e) => {
            e.composedPath();
            if (e.target.matches("[category-button]")) {
                button.innerHTML = "";
                button.innerHTML = await this.getHtml();
                url.searchParams.set('category', e.target.value);
                getPosts(url);
                e.stopImmediatePropagation();
            } else if (e.target.matches("[logout-button]")) {
                document.cookie = 'session=; Max-Age=-1;';
                document.getElementById("h_posts").click();
                e.stopImmediatePropagation();
            }
        });
    }
}