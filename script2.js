const cartItemsList = document.getElementById('cartItems');
const emptyCartMessage = document.getElementById('emptyCartMessage');
const buyButton = document.getElementById('buyButton');

// Получаем список предметов из localStorage
const cartItems = JSON.parse(localStorage.getItem('cartItems')) || [];

// Создаем объект для подсчета количества предметов
const itemCounter = {};

// Подсчет количества предметов
cartItems.forEach(item => {
  if (item.title in itemCounter) {
    itemCounter[item.title]++;
  } else {
    itemCounter[item.title] = 1;
  }
});

// Отображаем предметы в корзине с указанием количества
if (cartItems.length > 0) {
  for (const [title, count] of Object.entries(itemCounter)) {
    const cartItem = document.createElement('li');
    const item = cartItems.find(item => item.title === title);
    if (item) {
      cartItem.textContent = `${title} x${count} - ${item.price}`;
      cartItemsList.appendChild(cartItem);
    }
  }
}

// Очистка корзины при выходе
window.addEventListener('beforeunload', () => {
  localStorage.removeItem('cartItems');
});

// Проверка наличия предметов в корзине и отображение кнопки "Купить"
if (cartItems.length > 0) {
  buyButton.style.display = 'block';
}
