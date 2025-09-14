const API_BASE_URL = window.API_BASE_URL || 'http://localhost:8080';

document.getElementById('urlForm').addEventListener('submit', async function(e) {
    e.preventDefault();
    
    const longURL = document.getElementById('longURL').value;
    const submitBtn = document.getElementById('submitBtn');
    const loading = document.getElementById('loading');
    const result = document.getElementById('result');
    const error = document.getElementById('error');
    
    // Reset states
    result.classList.remove('show');
    error.classList.remove('show');
    loading.style.display = 'block';
    submitBtn.disabled = true;
    submitBtn.textContent = 'Encurtando...';
    
    try {
        const formData = new FormData();
        formData.append('longURL', longURL);
        
        const response = await fetch(API_BASE_URL + '/', {
            method: 'POST',
            body: formData
        });
        
        if (!response.ok) {
            throw new Error(`Erro ${response.status}: ${response.statusText}`);
        }
        
        const data = await response.json();
        
        // Display result
        document.getElementById('originalUrl').textContent = data.longURL;
        document.getElementById('shortUrl').textContent = `${API_BASE_URL}/${data.shortURL}`;
        result.classList.add('show');
        
    } catch (err) {
        console.error('Erro ao encurtar URL:', err);
        document.getElementById('errorMessage').textContent = err.message;
        error.classList.add('show');
    } finally {
        loading.style.display = 'none';
        submitBtn.disabled = false;
        submitBtn.textContent = 'Encurtar URL';
    }
});

function copyToClipboard(buttonEl) {
    const shortUrl = document.getElementById('shortUrl').textContent.trim();

    const setCopiedState = () => {
        const originalText = buttonEl.dataset.originalText || buttonEl.textContent;
        if (!buttonEl.dataset.originalText) {
            buttonEl.dataset.originalText = originalText;
        }
        buttonEl.textContent = 'Copiado';
        buttonEl.classList.add('copied');
        setTimeout(() => {
            buttonEl.textContent = buttonEl.dataset.originalText;
            buttonEl.classList.remove('copied');
        }, 1800);
    };

    const copyWithClipboardApi = () => navigator.clipboard.writeText(shortUrl);

    const copyWithFallback = () => {
        return new Promise((resolve, reject) => {
            const ta = document.createElement('textarea');
            ta.value = shortUrl;
            ta.setAttribute('readonly', '');
            ta.style.position = 'fixed';
            ta.style.top = '-9999px';
            document.body.appendChild(ta);
            ta.select();
            try {
                const successful = document.execCommand('copy');
                document.body.removeChild(ta);
                successful ? resolve() : reject(new Error('execCommand falhou'));
            } catch (e) {
                document.body.removeChild(ta);
                reject(e);
            }
        });
    };

    const isLocalhost = ['localhost', '127.0.0.1', '::1'].includes(location.hostname);
    const canUseClipboard = !!(navigator.clipboard && (window.isSecureContext || isLocalhost));

    (canUseClipboard ? copyWithClipboardApi() : copyWithFallback())
        .then(setCopiedState)
        .catch(err => {
            console.error('Erro ao copiar:', err);
            copyWithFallback()
                .then(setCopiedState)
                .catch(() => {
                    alert('Não foi possível copiar automaticamente. Copie manualmente: ' + shortUrl);
                });
        });
}
