import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("Posts");
    }

    async getHtml() {
        return `
        <h1>Welcome back, Dom</h1>
        <p>
            Fugiat voluptate et nisi
            </p>
            <p>
            <a href="/add' data-link>View recent posts</a>.
            </p>
        `;
    }
}