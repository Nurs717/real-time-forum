import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("Posts");
    }

    async getHtml() {
        return `
        <h1>Posts</h1>
        <p>
        Posts
        </p>
        <div id="posts"></div>
        <p>
            <a href="/add" data-link>Create Post</a>.
        </p>
        `;
    }

    async getPosts() {
        const url = "http://localhost:8080/"

        fetch(url, {
            method: "GET",
        }).then(
            response => response.json()
        ).then(
            (data) => {
                console.log('hello', data)
                document.getElementById("posts").innerHTML = data[1].post;
            }
        )
    }
}