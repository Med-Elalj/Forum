// Function to fetch random articles from Dev.to API
async function getRandomArticles(count = 5) {
    try {
        // Fetch articles from Dev.to API
        const response = await fetch(`https://dev.to/api/articles?per_page=${count}`);
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const articles = await response.json();
        
        // Process and format the articles
        const formattedArticles = articles.map(article => ({
            title: article.title,
            author: article.user.name,
            url: article.url,
            description: article.description,
            publishedDate: new Date(article.published_at).toLocaleDateString(),
            readingTime: article.reading_time_minutes,
            tags: article.tags
        }));

        return formattedArticles;
        
    } catch (error) {
        console.error('Error fetching articles:', error);
        throw error;
    }
}

// Example usage with DOM manipulation
function displayArticles() {

    getRandomArticles()
        .then(articles => {
            
            articles.forEach(article => {

                //console.log(articles);
                

             });
        })
        .catch(error => {
        });
}

// Add some basic CSS styles
const styles = `
    .article {
        margin-bottom: 20px;
        padding: 15px;
        border: 1px solid #ddd;
        border-radius: 5px;
    }
    .article h2 {
        margin-top: 0;
    }
    .article a {
        color: #2563eb;
        text-decoration: none;
    }
    .author {
        color: #666;
        font-style: italic;
    }
    .meta {
        font-size: 0.9em;
        color: #666;
        margin: 10px 0;
    }
    .meta span {
        margin-right: 15px;
    }
    .tags {
        display: flex;
        flex-wrap: wrap;
        gap: 5px;
    }
    .tag {
        background: #e5e7eb;
        padding: 3px 8px;
        border-radius: 15px;
        font-size: 0.8em;
    }
    .error {
        color: #dc2626;
        padding: 10px;
        border: 1px solid #dc2626;
        border-radius: 5px;
    }
`;

displayArticles()