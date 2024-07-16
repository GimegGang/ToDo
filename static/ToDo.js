document.addEventListener('DOMContentLoaded', function () {
    document.querySelectorAll('.edit').forEach(function (editButton) {
        editButton.addEventListener('click', function () {
            const todoDiv = this.closest('.todo');

            const titleElement = todoDiv.querySelector('h2');
            const descElement = todoDiv.querySelector('p');

            const currentTitle = titleElement.textContent;
            const currentDesc = descElement.textContent;

            const titleInput = document.createElement('input');
            titleInput.className = "editInp"
            titleInput.name = 'title';
            titleInput.value = currentTitle;

            const descInput = document.createElement('input');
            descInput.className = "editInp"
            descInput.name = 'desc';
            descInput.value = currentDesc;

            const saveButton = document.createElement('button');
            saveButton.className = "editB"
            saveButton.type = 'submit';
            saveButton.textContent = 'Save';

            const form = document.createElement('form');
            form.method = 'POST';
            form.action = `/edit/${todoDiv.dataset.id}`;

            form.appendChild(titleInput);
            form.appendChild(descInput);
            form.appendChild(saveButton);

            todoDiv.replaceChild(form, titleElement);
            todoDiv.replaceChild(form, descElement);
        });
    });
});
