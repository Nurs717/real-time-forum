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

    async Init() {
        this.getPosts();
    }

    async getPosts() {
        const url = "http://localhost:8080/"

        fetch(url, {
            method: "GET",
            credentials: 'include'
        }).then(
            (response) => {
                return response.json()
            }
        ).then(
            (data) => {
                // console.log('hello', data);
                let posts = document.getElementById("posts");
                data.map(post => {
                    let line = document.createElement('p');
                    line.innerText = post.post;
                    posts.appendChild(line);
                })
            }
        )
    }
}