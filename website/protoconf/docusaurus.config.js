// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'protoconf',
  tagline: 'Streamlining Configurations, Accelerating Deployments.',
  favicon: 'img/favicon.ico',

  markdown: {
    mermaid: true,
  },
  themes: ["@docusaurus/theme-mermaid"],
  // Set the production url of your site here
  url: 'https://www.protoconf.dev',
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: '/',

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: 'protoconf', // Usually your GitHub org/user name.
  projectName: 'www.protoconf.dev', // Usually your repo name.
  trailingSlash: false,
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        gtag: {
          trackingID: "G-H3NLBNNQD7"
        },
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/protoconf/protoconf/tree/main/website/protoconf',
          lastVersion: 'current',
          versions: {
            current: {
              label: 'v0.1.7',
              path: '0.1.7',
            }
          }
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/protoconf/protoconf/tree/main/website/protoconf',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      // Replace with your project's social card
      image: 'img/protoconf-social-card.png',
      navbar: {
        // title: 'protoconf',
        logo: {
          alt: 'Protoconf Logo',
          src: 'img/protoconf.png',
          srcDark: 'img/protoconf_inverse.png',
        },
        items: [
          {
            type: 'doc',
            docId: 'intro',
            position: 'left',
            label: 'Docs',
          },
          {
            type: 'dropdown',
            label: 'Integrations',
            position: 'left',
            items: [
              {
                label: 'Terraform',
                type: 'doc',
                docId: 'integrations/terraform'
              },
              {
                label: 'Envoy',
                type: 'doc',
                docId: 'integrations/envoy'
              }
            ]
          },
          {
            type: 'dropdown',
            label: 'Use Cases',
            position: 'left',
            items: [
              {
                label: 'Employee directory as Code',
                type: 'doc',
                docId: 'use-cases/employee-directory'
              },{
                label: 'Platform Engineering',
                type: 'doc',
                docId: 'use-cases/platform-engineering'
              }
            ]
          },
          {
            to: '/blog',
            position: 'left',
            label: 'Blog',
          },
          {
            type: 'docsVersionDropdown',
            position: 'right'
          },
          {
            href: 'https://github.com/protoconf/protoconf',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Docs',
            items: [
              {
                label: 'Tutorial',
                to: 'docs/0.1.7/intro',
              },
            ],
          },
          {
            title: 'Community',
            items: [
              {
                label: 'Discord',
                href: 'https://discord.protoconf.sh',
              },
              {
                label: 'Twitter',
                href: 'https://twitter.com/protoconfdev',
              },
            ],
          },
          {
            title: 'More',
            items: [
              {
                label: 'Blog',
                to: '/blog',
              },
              {
                label: 'GitHub',
                href: 'https://github.com/protoconf/protoconf',
              },
            ],
          },
        ],
        logo: {
          alt: 'Protoconf',
          src: 'img/protoconf_bottom.svg',
          href: 'https://www.protoconf.dev',
          width: 51,
          height: 51,
        },
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
        additionalLanguages: ['protobuf', 'rust', 'java', 'docker' ],
      },
    }),
};

module.exports = config;
