export async function drawNavMenuLoggedIn() {
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

export async function drawNavMenuLoggedOut() {
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

export async function drawWelcomeUserName() {
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

export const getError500 = `
    <div id="error-container">
        <div class="fof">
            <h1>Error 500</h1>
        </div>
    </div>
`

export const getError404 = `
    <div id="error-container">
        <div class="fof">
            <h1>Error 404</h1>
        </div>
    </div>
`