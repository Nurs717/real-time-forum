import AbstractView from "./AbstractView.js";

async function getPosts(url) {
    // document.cookie = 'session=; Max-Age=-1;';
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
                //draw categories
                drawCategories();
            } else {
                // draw nav menu logged out
                drawNavmenuLoggedOut();
                //draw categories
                drawCategories();
            }
            return response.json()
        }
    ).then(
        (data) => {
            console.log('hello', data);
            let app = document.getElementById("app");
            let username = document.getElementById('welcome_username');
            username.innerHTML = data[0].username;
            let posts = document.createElement("div");
            posts.setAttribute('id', 'posts');
            app.appendChild(posts);
            data.map(post => {
                let line = document.createElement('p');
                line.innerText = 'username: ' + post.username + ' title: ' + post.post_title;
                posts.appendChild(line);
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
    logout.setAttribute('my-button', '');
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
    sport.setAttribute('my-button', '');
    sport.innerHTML = 'sport';
    let religion = document.createElement('button');
    religion.setAttribute('value', 'religion');
    religion.setAttribute('my-button', '');
    religion.innerHTML = 'religion';
    let programming = document.createElement('button');
    programming.setAttribute('my-button', '');
    programming.setAttribute('value', 'programming');
    programming.innerHTML = 'programming'
    categories.appendChild(sport);
    categories.appendChild(religion);
    categories.appendChild(programming);
    row[1].appendChild(categories);
}

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("MyForum");
        this.myUrl = "http://localhost:8080/";
    }

    async getHtml() {
        return `
        <header id="#top">
            <div class="row">
                <div class="column lpad">
                    <div class="logo">
                        <span>MyForum</span>
                    </div>
                </div>
                <div class="column ar lpad">
                    <nav class="menu">
                        <a href="/" class="current nav_link" data-link>Posts</a>
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

        const button = document.getElementById('app');
        button.addEventListener("click", async(e) => {
            e.preventDefault();
            if (e.target.matches("[my-button]")) {
                button.innerHTML = "";
                button.innerHTML = await this.getHtml();
                url.searchParams.set('category', e.target.value);
                getPosts(url);
                console.log(url.href)
            }
        });
    }
}