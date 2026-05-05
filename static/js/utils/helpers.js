export const debounce = (func, wait) => {
  let timeout;
  return function executedFunction(...args) {
    const later = () => {
      clearTimeout(timeout);
      func(...args);
    };
    clearTimeout(timeout);
    timeout = setTimeout(later, wait);
  };
};

export const throttle = (func, limit) => {
  let inThrottle;
  return function executedFunction(...args) {
    if (!inThrottle) {
      func(...args);
      inThrottle = true;
      setTimeout(() => (inThrottle = false), limit);
    }
  };
};

export const formatEmail = (email) => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(email);
};

export const validateForm = (form) => {
  const requiredFields = form.querySelectorAll('[required]');
  let isValid = true;

  requiredFields.forEach(field => {
    if (!field.value.trim()) {
      isValid = false;
      showFieldError(field);
    } else {
      hideFieldError(field);
      if (field.type === 'email' && !formatEmail(field.value)) {
        isValid = false;
        showFieldError(field, '请输入有效的邮箱地址');
      }
    }
  });

  return isValid;
};

export const showFieldError = (field, message = '此字段为必填项') => {
  field.classList.add('border-red-500');
  const errorElement = field.parentElement.querySelector('.error-message');
  if (errorElement) {
    errorElement.textContent = message;
    errorElement.classList.remove('hidden');
  }
};

export const hideFieldError = (field) => {
  field.classList.remove('border-red-500');
  const errorElement = field.parentElement.querySelector('.error-message');
  if (errorElement) {
    errorElement.classList.add('hidden');
  }
};

export const scrollToSection = (sectionId, offset = 80) => {
  const element = document.querySelector(sectionId);
  if (element) {
    window.scrollTo({
      top: element.offsetTop - offset,
      behavior: 'smooth'
    });
  }
};

export const addClassOnScroll = (element, className, threshold = 50) => {
  window.addEventListener('scroll', throttle(() => {
    if (window.scrollY > threshold) {
      element.classList.add(className);
    } else {
      element.classList.remove(className);
    }
  }, 100));
};

export const animateOnScroll = (selector, animationClass) => {
  if ('IntersectionObserver' in window) {
    const observer = new IntersectionObserver((entries) => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          entry.target.classList.add(animationClass);
        }
      });
    }, {
      rootMargin: '0px 0px -10% 0px',
      threshold: 0.1
    });

    document.querySelectorAll(selector).forEach(el => observer.observe(el));
  }
};

export const createModal = (options) => {
  const { title, message, type, onClose } = options;
  
  const modal = document.createElement('div');
  modal.className = 'fixed inset-0 z-50 flex items-center justify-center bg-black/50';
  
  modal.innerHTML = `
    <div class="bg-white rounded-xl p-8 shadow-2xl max-w-md w-full mx-4 animate-fade-in">
      <div class="text-center">
        <div class="w-16 h-16 mx-auto mb-4 rounded-full flex items-center justify-center ${
          type === 'success' ? 'bg-green-100' : 'bg-red-100'
        }">
          <i class="fa ${type === 'success' ? 'fa-check' : 'fa-times'} text-2xl ${
            type === 'success' ? 'text-green-600' : 'text-red-600'
          }"></i>
        </div>
        <h3 class="text-xl font-bold mb-2 ${
          type === 'success' ? 'text-green-600' : 'text-red-600'
        }">${title}</h3>
        <p class="text-muted mb-6">${message}</p>
        <button id="modal-close" class="w-full bg-primary hover:bg-primary/90 text-white py-3 rounded-lg transition-colors">
          关闭
        </button>
      </div>
    </div>
  `;

  document.body.appendChild(modal);

  const closeModal = () => {
    modal.remove();
    if (typeof onClose === 'function') {
      onClose();
    }
  };

  modal.addEventListener('click', (e) => {
    if (e.target === modal || e.target.id === 'modal-close') {
      closeModal();
    }
  });

  return modal;
};

export const formatPhone = (phone) => {
  return phone.replace(/(\d{3})(\d{4})(\d{4})/, '$1-$2-$3');
};

export const getRandomInt = (min, max) => {
  return Math.floor(Math.random() * (max - min + 1)) + min;
};