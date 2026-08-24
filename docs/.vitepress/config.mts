import { defineConfig } from 'vitepress'
export default defineConfig({
  appearance: 'dark',
  base: '/uast/',
  head: [
    ['meta', { name: 'viewport', content: 'width=device-width, initial-scale=1.0, viewport-fit=cover' }]
  ],
  cleanUrls: true,
  lastUpdated: true,
  locales: {
    en: {
      description: 'A high-performance, zero-dependency type‑safe query builder.',
      label: 'English',
      lang: 'en',
      link: '/en/',
      title: 'UAST',
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
            text: 'License', 
            link: '/en/license' 
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
                items: [
                  { 
                    text: 'Core',
                    items: [
                      { 
                        text: 'Constructors', 
                        link: '/en/core_constructors' 
                      },
                      { 
                        text: 'Methods', 
                        link: '/en/core_methods' 
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
                  },
                  { 
                    text: 'SQL',
                    items: [
                      { 
                        text: 'Constructors', 
                        link: '/en/sql_constructors' 
                      },
                      { 
                        text: 'Methods', 
                        link: '/en/sql_methods' 
                      },
                      { 
                        text: 'Options', 
                        link: '/en/sql_options' 
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
          message: 'Dual-licensed under AGPLv3 and Commercial',
          copyright: '© 2026 Mikhail Dadaev'
        }
      }
    },
    ru: {
      description: 'Производительный и типобезопасный SQL-билдер с нулевыми аллокациями..',
      label: 'Русский',
      lang: 'ru',
      link: '/ru/',
      title: 'UAST',
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
            text: 'Лицензия', 
            link: '/ru/license' 
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
          }
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
                items: [
                  { 
                    text: 'Ядро',
                    items: [
                      { 
                        text: 'Конструкторы', 
                        link: '/ru/core_constructors' 
                      },
                      { 
                        text: 'Методы', 
                        link: '/ru/core_methods' 
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
                  },
                  { 
                    text: 'SQL',
                    items: [
                      { 
                        text: 'Конструкторы', 
                        link: '/ru/sql_constructors' 
                      },
                      { 
                        text: 'Методы', 
                        link: '/ru/sql_methods' 
                      },
                      { 
                        text: 'Опции', 
                        link: '/ru/sql_options' 
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
          message: 'Двойное лицензирование: AGPLv3 и Коммерческая',
          copyright: '© 2026 Михаил Дадаев'
        }
      }
    },
    zh: {
      description: '高性能、零分配、类型安全的 SQL 构建器。',
      label: '简体中文',
      lang: 'zh',
      link: '/zh/',
      title: 'UAST',
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
            text: '许可证', 
            link: '/zh/license' 
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
          }
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
                items: [
                  { 
                    text: '核心',
                    items: [
                      { 
                        text: '构造函数', 
                        link: '/zh/core_constructors' 
                      },
                      { 
                        text: '方法', 
                        link: '/zh/core_methods' 
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
                  },
                  { 
                    text: 'SQL',
                    items: [
                      { 
                        text: '构造函数', 
                        link: '/zh/sql_constructors' 
                      },
                      { 
                        text: '方法', 
                        link: '/zh/sql_methods' 
                      },
                      { 
                        text: '选项', 
                        link: '/zh/sql_options' 
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
          message: '根据 AGPLv3 和商业许可证双许可',
          copyright: '© 2026 Mikhail Dadaev'
        }
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