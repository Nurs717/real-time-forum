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
    }
}