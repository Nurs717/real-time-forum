import AbstractView from "./AbstractView.js";
import {drawNavMenuLoggedIn, drawWelcomeUserName, drawNavMenuLoggedOut, getError500} from "./Shared.js";

async function getPosts(url) {
    fetch(url.href, {
        method: "GET",
        credentials: 'include'
    }).then(
        (response) => {
            if (response.status === 500) {
                let app = document.getElementById("app");
                app.innerHTML = getError500;
                return
            } else if (response.ok) {
                // draw nav menu
                drawNavMenuLoggedIn();
                //draw username
                drawWelcomeUserName();
            } else if (response.status === 403) {
                // draw nav menu logged out
                drawNavMenuLoggedOut();
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
            let t_body = document.getElementById("tbody");
            if (document.cookie !== "") {
                let username = document.getElementById('welcome_username');
                username.innerHTML = data.username;
            }
            data?.posts?.map(post => {
                let tr = document.createElement("tr");
                if (post.ID % 2 === 0) {
                    tr.setAttribute('class', 'even');
                } else {
                    tr.setAttribute('class', 'odd');
                }
                let td_topic = document.createElement("td");
                td_topic.style.cssText = "padding-left: 25px; padding-right: 15px; text-align: justify;"
                let a_topic = document.createElement("a");
                a_topic.setAttribute("class", "a_topic")
                a_topic.setAttribute("href", `http://localhost:8081/post/${post.ID}`);
                a_topic.setAttribute("data-link", "")
                a_topic.innerText = post.post_title;
                td_topic.appendChild(a_topic);
                let td_categories = document.createElement("td");
                let div_categories = document.createElement("div");
                let cat = '';
                if  (post.category !== null || post.category !== 'undefined') {
                    cat = post.category.join(", ");
                }
                div_categories.innerText = cat;
                td_categories.appendChild(div_categories);
                let td_comments = document.createElement("td");
                let div_comments = document.createElement("div");
                div_comments.innerText = '2';
                td_comments.appendChild(div_comments);
                let td_created = document.createElement("td");
                let div_created = document.createElement("div");
                div_created.innerText = `by ${post.username}\n${post.post_date}`;
                td_created.appendChild(div_created);
                tr.appendChild(td_topic);
                tr.appendChild(td_categories);
                tr.appendChild(td_comments);
                tr.appendChild(td_created);
                t_body.appendChild(tr);
            })
        }
    )
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
    programming.innerHTML = 'programming';
    let all = document.createElement('button');
    all.setAttribute('category-button', '');
    all.setAttribute('value', '');
    all.innerHTML = 'all';
    categories.appendChild(sport);
    categories.appendChild(religion);
    categories.appendChild(programming);
    categories.appendChild(all);
    row[1].appendChild(categories);
}

async function drawPosts() {
    let app = document.getElementById("app");
    let posts = document.createElement("div");
    posts.setAttribute('id', 'posts');
    posts.setAttribute('class', 'row mt');
    let table = document.createElement("table");
    table.setAttribute('class', 'table');
    let table_head = document.createElement("thead");
    let tr_head = document.createElement("tr");
    let th_topic = document.createElement("th");
    th_topic.style.cssText = "width:60%";
    th_topic.innerHTML = "Topic";
    let th_comments = document.createElement("th");
    th_comments.style.cssText = "width:10%";
    th_comments.innerHTML = "Comments";
    let th_created = document.createElement("th");
    th_created.style.cssText = "width:15%";
    th_created.innerHTML = "Created";
    let th_categories = document.createElement("th");
    th_categories.style.cssText = "width:15%"
    th_categories.innerHTML = "Categories";
    tr_head.appendChild(th_topic);
    tr_head.appendChild(th_categories);
    tr_head.appendChild(th_comments);
    tr_head.appendChild(th_created);
    table_head.appendChild(tr_head);
    let table_body = document.createElement("tbody");
    table_body.setAttribute('id', 'tbody')
    table.appendChild(table_head);
    table.appendChild(table_body);
    posts.appendChild(table);
    app.appendChild(posts);
}

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
        let url = new URL("http://localhost:8080/posts");
        await getPosts(url);

        const button = document.getElementById("app");
        button.addEventListener("click", async(e) => {
            e.composedPath();
            if (e.target.matches("[category-button]")) {
                button.innerHTML = "";
                button.innerHTML = await this.getHtml();
                url.searchParams.set('category', e.target.value);
                await getPosts(url);
                e.stopImmediatePropagation();
            } else if (e.target.matches("[logout-button]")) {
                document.cookie = 'session=; Max-Age=-1;';
                document.getElementById("h_posts").click();
                e.stopImmediatePropagation();
            }
        });
    }
}