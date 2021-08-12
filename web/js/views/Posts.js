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
            Here will be Posts
        </p>
        <p>
            <a href="/add" data-link>Create Post</a>.
        </p>
        `;
    }
}