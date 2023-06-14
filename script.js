
    const cartButton = document.getElementById('cartButton');
    const addToCartButtons = document.querySelectorAll('.add-to-cart');
    const itemCount = document.querySelector('.item-count');
    
    // Получаем список предметов из localStorage
    let cartItems = JSON.parse(localStorage.getItem('cartItems')) || [];
    
    // Устанавливаем начальное значение счетчика
    itemCount.textContent = cartItems.length;
    
    // Функция для перенаправления на страницу "cart.html"
    function redirectToCart() {
        window.location.href = 'cart.html'; // Замените 'cart.html' на URL вашей страницы корзины
    }
    
    // Добавляем обработчик события клика на кнопку корзины
    cartButton.addEventListener('click', redirectToCart);
    
    // Функция, которая добавляет предмет в корзину
    function addToCart(event) {
        const item = event.target.parentNode;
        const title = item.querySelector('.menu__title').textContent;
        const price = item.querySelector('.menu__price').textContent;
    
        // Создаем объект для предмета
        const newItem = {
            title: title,
            price: price
        };
    
        // Получаем текущий список предметов в корзине из localStorage
        cartItems = JSON.parse(localStorage.getItem('cartItems')) || [];
    
        // Добавляем новый предмет в список
        cartItems.push(newItem);
    
        // Обновляем количество предметов в счетчике
        itemCount.textContent = cartItems.length;
    
        // Показываем счетчик, если он был скрыт
        itemCount.style.display = 'block';
    
        // Добавляем класс для отображения красного круга
        cartButton.classList.add('active');
    
        // Сохраняем обновленный список в localStorage
        localStorage.setItem('cartItems', JSON.stringify(cartItems));
    }
    
    // Добавляем обработчик события клика на каждую кнопку "Add to cart"
    addToCartButtons.forEach(button => {
        button.addEventListener('click', addToCart);
    });
