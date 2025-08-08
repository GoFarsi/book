document.addEventListener('DOMContentLoaded', () => {
  document.querySelectorAll('.go-playground').forEach(container => {
    const runBtn = container.querySelector('.run-code');
    const copyBtn = container.querySelector('.copy-code');
    const codeBlock = container.querySelector('code');
    const outputBlock = container.querySelector('.run-output');

     // Copy logic
    copyBtn.addEventListener('click', async () => {
      try {
        await navigator.clipboard.writeText(codeBlock.innerText.trim());
        copyBtn.textContent = 'کپی شد ✅';
        setTimeout(() => {
          copyBtn.textContent = 'کپی';
        }, 1500);
      } catch (err) {
        console.error('Copy failed:', err);
        copyBtn.textContent = '❌ خطا';
      }
    });

    runBtn.addEventListener('click', async () => {
      const rawCode = codeBlock.innerText.trim();
      outputBlock.classList.remove('hidden');

      if (!rawCode) {
        outputBlock.textContent = '⚠️ کدی برای اجرا وارد نشده است.';
        return;
      }

      outputBlock.textContent = '⏳ درحال فرمت و اجرا...';

      try {
        const formattedCode = await formatGoCode(rawCode);

        const runResp = await fetch('http://localhost:9123/api/v2/run?vet=true&backend=', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            files: { 'main.go': formattedCode }
          })
        });

        const runResult = await runResp.json();

          if (runResult.errors) {
            outputBlock.textContent = `${runResult.errors}`;
          } else if (Array.isArray(runResult.events)) {
            outputBlock.textContent = `${runResult.events.map(e => e.Message).join('')}`;
          } else {
            outputBlock.textContent = '⚠️ اجرای کد انجام نشد. پاسخ نامعتبر بود.';
          }

      } catch (err) {
        outputBlock.textContent = `${err.message}`;
        console.error(err);
      }
    });
  });
});

async function formatGoCode(code) {
  const response = await fetch('http://localhost:9123/api/v2/format?backend=', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      files: { 'main.go': code }
    })
  });

  const result = await response.json();

  if (!response.ok) {
    throw new Error(result.error || 'خطای ناشناخته در فرمت کردن کد');
  }

  return result.files['main.go'];
}

// Manually re-highlight all Prism blocks after DOM is ready
document.addEventListener('DOMContentLoaded', () => {
  Prism.highlightAll();
});
