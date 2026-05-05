import { defineConfig } from 'vitepress'
export default defineConfig({
  appearance: 'dark',
  base: '/uast/',
  head: [
    ['link', { rel: 'stylesheet', href: '/uast/styles.css' }],
    ['script', { src: '/uast/scripts.js' }]
  ],
  lastUpdated: true,
  locales: {
    en: {
      description: 'A high-performance, zero‑allocation type‑safe SQL builder.',
      label: 'English',
      lang: 'en',
      link: '/en/',
      title: 'uast',
      themeConfig: {
        search: {
          provider: 'local',
          options: {
            translations: {
              button: {
                buttonText: 'Search',
                buttonAriaLabel: 'Search'
              },
              modal: {
                noResultsText: 'No results found',
                resetButtonTitle: 'Clear search',
                footer: {
                  selectText: 'to select',
                  navigateText: 'to navigate',
                  closeText: 'to close'
                }
              }
            }
          }
        },
        nav: [
          { 
            text: 'Home', 
            link: '/en/' 
          },
          { 
            text: 'Go', 
            link: '/en/go' 
          },
          { 
            text: 'Benchmarks', 
            link: '/en/benchmarks' 
          },
          { 
            text: 'API', 
            link: '/en/core_constructors' 
          },
        ],
        sidebar: [
          {
            items: [
              { 
                text: 'Go', 
                link: '/en/go' 
              },
              { 
                text: 'Benchmarks', 
                link: '/en/benchmarks' 
              },
              { 
                text: 'API', 
                collapsed: true,
                items: [
                  { 
                    text: 'Core', 
                    collapsed: true,
                    items: [
                      { 
                        text: 'Constructors', 
                        link: '/en/core_constructors' 
                      },
                      { 
                        text: 'Options', 
                        link: '/en/core_options' 
                      },
                      { 
                        text: 'Types', 
                        link: '/en/core_types' 
                      }
                    ] 
                  }
                ] 
              }
            ]
          }
        ],
        darkModeSwitchLabel: "Appearance",
        darkModeSwitchTitle: "Switch to dark theme",
        lightModeSwitchTitle: "Switch to light theme",
        sidebarMenuLabel: "Menu",
        returnToTopLabel: "Return to top",
        outline: {
          label: "On this page"
        },
        lastUpdated: {
          text: "Last Updated",
          formatOptions: {
            dateStyle: "short",
            timeStyle: "short"
          }
        },
        docFooter: {
          prev: "Previous page",
          next: "Next page"
        },
        footer: {
          message: 'Released under the Apache License 2.0',
          copyright: '© 2026 Mikhail Dadaev'
        }
      }
    },
    ru: {
      description: 'Производительный и типобезопасный SQL-билдер с нулевыми аллокациями..',
      label: 'Русский',
      lang: 'ru',
      link: '/ru/',
      title: 'uast',
      themeConfig: {
        search: {
          provider: 'local',
          options: {
            translations: {
              button: {
                buttonText: 'Поиск',
                buttonAriaLabel: 'Поиск'
              },
              modal: {
                noResultsText: 'Ничего не найдено',
                resetButtonTitle: 'Очистить поиск',
                footer: {
                  selectText: 'выбрать',
                  navigateText: 'перейти',
                  closeText: 'закрыть'
                }
              }
            }
          }
        },
        nav: [
          { 
            text: 'Главная', 
            link: '/ru/' 
          },
          { 
            text: 'Go', 
            link: '/ru/go' 
          },
          { 
            text: 'Бенчмарки', 
            link: '/ru/benchmarks' 
          },
          { 
            text: 'API', 
            link: '/ru/core_constructors' 
          },
        ],
        sidebar: [
          {
            items: [
              { 
                text: 'Go', 
                link: '/ru/go' 
              },
              { 
                text: 'Бенчмарки', 
                link: '/ru/benchmarks' 
              },
              {
                text: 'API', 
                collapsed: true,
                items: [
                  { 
                    text: 'Ядро', 
                    collapsed: true,
                    items: [
                      { 
                        text: 'Конструкторы', 
                        link: '/ru/core_constructors' 
                      },
                      { 
                        text: 'Опции', 
                        link: '/ru/core_options' 
                      },
                      { 
                        text: 'Типы', 
                        link: '/ru/core_types' 
                      }
                    ] 
                  }
                ]
              }
            ]
          }
        ],
        darkModeSwitchLabel: "Внешний вид",
        darkModeSwitchTitle: "Переключиться на тёмную тему",
        lightModeSwitchTitle: "Переключиться на светлую тему",
        sidebarMenuLabel: "Меню",
        returnToTopLabel: "Вернуться наверх",
        outline: {
          label: "Содержание страницы"
        },
        lastUpdated: {
          text: "Последние изменения",
          formatOptions: {
            dateStyle: "short",
            timeStyle: "short"
          }
        },
        docFooter: {
          prev: "Предыдущая страница",
          next: "Следующая страница"
        },
        footer: {
          message: 'Под лицензией Apache 2.0',
          copyright: '© 2026 Михаил Дадаев'
        },
      }
    },
    zh: {
      description: '高性能、零分配、类型安全的 SQL 构建器。',
      label: '简体中文',
      lang: 'zh',
      link: '/zh/',
      title: 'uast',
      themeConfig: {
        search: {
          provider: 'local',
          options: {
            translations: {
              button: {
                buttonText: '搜索',
                buttonAriaLabel: '搜索'
              },
              modal: {
                noResultsText: '未找到结果',
                resetButtonTitle: '清除搜索',
                footer: {
                  selectText: '选择',
                  navigateText: '导航',
                  closeText: '关闭'
                }
              }
            }
          }
        },
        nav: [
          { 
            text: '首页', 
            link: '/zh/' 
          },
          { 
            text: 'Go', 
            link: '/zh/go' 
          },
          { 
            text: '基准测试', 
            link: '/zh/benchmarks' 
          },
          { 
            text: 'API', 
            link: '/zh/core_constructors' 
          },
        ],
        sidebar: [
          {
            items: [
              { 
                text: 'Go', 
                link: '/zh/go' 
              },
              { 
                text: '基准', 
                link: '/zh/benchmarks' 
              },
              { 
                text: 'API', 
                collapsed: true,
                items: [
                  { 
                    text: '核心', 
                    collapsed: true,
                    items: [
                      { 
                        text: '构造函数', 
                        link: '/zh/core_constructors' 
                      },
                      { 
                        text: '选项', 
                        link: '/zh/core_options' 
                      },
                      { 
                        text: '类别', 
                        link: '/zh/core_types' 
                      }
                    ] 
                  }
                ] 
              }
            ]
          }
        ],
        darkModeSwitchLabel: "深色模式",
        darkModeSwitchTitle: "切换至深色主题",
        lightModeSwitchTitle: "切换至浅色主题",
        sidebarMenuLabel: "目录",
        returnToTopLabel: "返回至顶部",
        outline: {
          label: "页面导航"
        },
        lastUpdated: {
          text: "最近更改",
          formatOptions: {
            dateStyle: "short",
            timeStyle: "short"
          }
        },
        docFooter: {
          prev: "上一页",
          next: "下一页"
        },
        footer: {
          message: '根据 Apache 2.0 许可证发布',
          copyright: '© 2026 Mikhail Dadaev'
        },
      }
    }
  },
  themeConfig: {
    search: {
      provider: 'local'
    },
    socialLinks: [
      { 
        icon: 'github', 
        link: 'https://github.com/mikhaildadaev/uast' 
      }
    ],
  }
})