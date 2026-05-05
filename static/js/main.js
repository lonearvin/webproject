class App {
  constructor() {
    this.init();
  }

  init() {
    this.setupNavigation();
    this.setupFormValidation();
    this.setupIntersectionObserver();
    this.setupCaseFilter();
    this.setupModalHandlers();
    this.setupSmoothScroll();
    this.setupLazyLoading();
  }

  setupNavigation() {
    const navbar = document.getElementById('navbar');
    const menuToggle = document.getElementById('menu-toggle');
    const mobileMenu = document.getElementById('mobile-menu');

    window.addEventListener('scroll', () => {
      if (window.scrollY > 50) {
        navbar.classList.add('bg-glass', 'shadow-md');
      } else {
        navbar.classList.remove('bg-glass', 'shadow-md');
      }
    });

    menuToggle?.addEventListener('click', () => {
      mobileMenu?.classList.toggle('hidden');
      const isExpanded = menuToggle.getAttribute('aria-expanded') === 'true';
      menuToggle.setAttribute('aria-expanded', !isExpanded);
      mobileMenu?.setAttribute('aria-hidden', isExpanded);
    });
  }

  setupFormValidation() {
    const contactForm = document.getElementById('contactForm');
    if (!contactForm) return;

    const requiredFields = contactForm.querySelectorAll('input[required], textarea[required]');
    const emailField = contactForm.querySelector('input[type="email"]');
    const emailFormatError = contactForm.querySelector('.email-format-error');

    requiredFields.forEach(field => {
      field.addEventListener('blur', () => this.validateField(field));
      field.addEventListener('focus', () => {
        const errorElement = field.parentElement.querySelector('.error-message');
        this.hideError(field, errorElement);
        if (field.type === 'email') {
          emailFormatError?.classList.add('hidden');
        }
      });
    });

    emailField?.addEventListener('input', () => {
      if (emailField.value.trim()) {
        emailFormatError?.classList.add('hidden');
      }
    });

    contactForm.addEventListener('submit', async (e) => {
      e.preventDefault();
      await this.handleFormSubmit(contactForm);
    });
  }

  validateField(field) {
    const value = field.value.trim();
    const errorElement = field.parentElement.querySelector('.error-message');

    if (!value) {
      this.showError(field, errorElement);
      return false;
    } else {
      this.hideError(field, errorElement);
      return true;
    }
  }

  validateEmailFormat(email) {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  }

  showError(field, errorElement) {
    field.classList.add('border-red-500');
    errorElement?.classList.remove('hidden');
  }

  hideError(field, errorElement) {
    field.classList.remove('border-red-500');
    errorElement?.classList.add('hidden');
  }

  async handleFormSubmit(form) {
    const submitButton = form.querySelector('button[type="submit"]');
    if (!submitButton) return;

    submitButton.disabled = true;
    submitButton.textContent = '提交中...';

    try {
      const response = await fetch(form.action, {
        method: form.method,
        body: new FormData(form),
      });

      if (response.ok) {
        this.showModal('提交成功', '消息已成功发送，我们会尽快回复！', 'success');
      } else if (response.status === 409) {
        const data = await response.json();
        this.showModal('提交成功', data.message || '消息已经提交过了！', 'success');
      } else {
        const data = await response.json();
        this.showModal('提交失败', data.message || '请检查输入内容！', 'error');
      }
    } catch (error) {
      this.showModal('网络错误', '请检查网络连接后重试！', 'error');
    } finally {
      this.resetForm(form);
    }
  }

  resetForm(form) {
    const submitButton = form.querySelector('button[type="submit"]');
    form.reset();
    submitButton.disabled = false;
    submitButton.textContent = '发送消息';
  }

  setupIntersectionObserver() {
    const observerOptions = {
      root: null,
      rootMargin: '0px 0px -10% 0px',
      threshold: 0.1
    };

    const observer = new IntersectionObserver((entries) => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          entry.target.classList.add('visible');
        }
      });
    }, observerOptions);

    document.querySelectorAll('.animate-fade-in, .animate-slide-left, .animate-slide-right').forEach(el => {
      observer.observe(el);
    });

    const serviceCards = document.querySelectorAll('[data-service-card]');
    serviceCards.forEach((card, index) => {
      card.style.transitionDelay = `${index * 100}ms`;
    });

    const caseCards = document.querySelectorAll('[data-case-card]');
    caseCards.forEach((card, index) => {
      card.style.transitionDelay = `${index * 150}ms`;
    });
  }

  setupCaseFilter() {
    const filterButtons = document.querySelectorAll('.case-filter');
    filterButtons.forEach(button => {
      button.addEventListener('click', () => {
        filterButtons.forEach(btn => {
          btn.classList.remove('active', 'bg-primary', 'text-white');
          btn.classList.add('bg-white', 'text-dark');
        });
        button.classList.add('active', 'bg-primary', 'text-white');
        button.classList.remove('bg-white', 'text-dark');
      });
    });
  }

  setupModalHandlers() {
    const subscribeForm = document.querySelector('form[action="/subscribe"]');
    subscribeForm?.addEventListener('submit', async (e) => {
      e.preventDefault();
      const submitButton = subscribeForm.querySelector('button[type="submit"]');
      if (!submitButton) return;

      submitButton.disabled = true;

      try {
        const response = await fetch(subscribeForm.action, {
          method: subscribeForm.method,
          body: new FormData(subscribeForm),
        });

        if (response.ok) {
          this.showModal('订阅成功', '恭喜订阅成功！', 'success');
        } else if (response.status === 409) {
          this.showModal('订阅成功', '您已经订阅过了。', 'success');
        } else {
          this.showModal('订阅失败', '订阅失败，请重试！', 'error');
        }
      } catch (error) {
        this.showModal('网络错误', '请检查网络连接后重试', 'error');
      } finally {
        const form = document.querySelector('form[action="/subscribe"]');
        form?.reset();
        submitButton.disabled = false;
      }
    });
  }

  showModal(title, message, type) {
    const modal = document.getElementById('messageModal');
    const content = document.getElementById('modalContent');

    content.innerHTML = `
      <h3 class="text-lg font-bold mb-2 ${type === 'success' ? 'text-green-600' : 'text-red-600'}">
        ${title}
      </h3>
      <p class="text-muted">${message}</p>
    `;

    modal?.classList.remove('hidden');

    const closeModal = () => {
      modal?.classList.add('hidden');
      document.getElementById('closeModal')?.removeEventListener('click', closeModal);
      modal?.removeEventListener('click', closeModal);
    };

    document.getElementById('closeModal')?.addEventListener('click', closeModal);
    modal?.addEventListener('click', (e) => {
      if (e.target === modal) {
        closeModal();
      }
    });
  }

  setupSmoothScroll() {
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
      anchor.addEventListener('click', function (e) {
        e.preventDefault();
        const targetId = this.getAttribute('href');
        if (targetId === '#') return;

        const targetElement = document.querySelector(targetId);
        if (targetElement) {
          window.scrollTo({
            top: targetElement.offsetTop - 80,
            behavior: 'smooth'
          });

          const mobileMenu = document.getElementById('mobile-menu');
          if (mobileMenu && !mobileMenu.classList.contains('hidden')) {
            mobileMenu.classList.add('hidden');
            const menuToggle = document.getElementById('menu-toggle');
            menuToggle?.setAttribute('aria-expanded', 'false');
          }
        }
      });
    });
  }

  setupLazyLoading() {
    const lazyImages = document.querySelectorAll('img[loading="lazy"]');
    
    if ('IntersectionObserver' in window) {
      const imageObserver = new IntersectionObserver((entries, observer) => {
        entries.forEach(entry => {
          if (entry.isIntersecting) {
            const img = entry.target;
            img.src = img.dataset.src || img.src;
            img.classList.remove('opacity-0');
            observer.unobserve(img);
          }
        });
      }, {
        rootMargin: '50px',
        threshold: 0.1
      });

      lazyImages.forEach(img => imageObserver.observe(img));
    }
  }
}

document.addEventListener('DOMContentLoaded', () => {
  new App();
});