document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('new-article-form');
    const articleList = document.getElementById('article-list');

    form.addEventListener('submit', async (event) => {
        event.preventDefault();
        const title = document.getElementById('title').value;
        const content = document.getElementById('content').value;
        const author = document.getElementById('author').value;

        const response = await fetch('/articles', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ title, content, author }),
        });

        if (response.ok) {
            loadArticles();
            form.reset();
        } else {
            alert('Failed to add article');
        }
    });

    async function loadArticles() {
        const response = await fetch('/articles');
        const data = await response.json();
        articleList.innerHTML = '';
        data.articles.forEach(article => {
            const li = document.createElement('li');
            li.textContent = `${article.title} by ${article.author}`;
            articleList.appendChild(li);
        });
    }

    loadArticles();
});
