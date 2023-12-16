import CodeViewer, { CodeSchema } from "@/components/code";
import { Themes } from "@/util/themes";
import { useTheme } from "next-themes";
import Image from "next/image";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";

export type MarkdownProps = {
  markdown: string;
};

const MarkdownViewer = (props: MarkdownProps) => {
  const { theme } = useTheme();

  return (
    <div
      className={
        theme === Themes.DARK ? "prose dark:prose-invert" : "prose prose-slate"
      }
    >
      <Markdown
        remarkPlugins={[remarkGfm]}
        components={{
          a(props) {
            const { href, children, ...rest } = props;
            const url = new URL(href || "");
            if (url.protocol !== "http:" && url.protocol !== "https:") {
              return <></>;
            }
            if (url.host === "sliver.sh") {
              return (
                <a
                  {...rest}
                  href={url.toString()}
                  className="text-primary hover:text-primary-dark"
                >
                  {children}
                </a>
              );
            }
            return (
              <a
                {...rest}
                href={url.toString()}
                rel="noopener noreferrer"
                target="_blank"
                className="text-primary hover:text-primary-dark"
              >
                {children}
              </a>
            );
          },

          pre(props) {
            // We need to look at the child nodes to avoid wrapping
            // a monaco code block in a <pre> tag
            const { children, className, node, ...rest } = props;
            const childClass = (children as any)?.props?.className;
            if (
              childClass &&
              childClass.startsWith("language-") &&
              childClass !== "language-plaintext"
            ) {
              // @ts-ignore
              return <div {...rest}>{children}</div>;
            }

            return (
              <pre {...rest} className={className}>
                {children}
              </pre>
            );
          },

          img(props) {
            const { src, alt, ...rest } = props;
            return (
              // @ts-ignore
              <Image
                {...rest}
                src={src || ""}
                alt={alt || ""}
                width={500}
                height={500}
                className="w-full rounded-medium"
              />
            );
          },

          code(props) {
            const { children, className, node, ...rest } = props;
            const langTag = /language-(\w+)/.exec(className || "");
            const lang = langTag ? langTag[1] : "plaintext";
            if (lang === "plaintext") {
              return (
                <span className="prose-sm">
                  <code {...rest} className={className}>
                    {children}
                  </code>
                </span>
              );
            }
            return (
              <CodeViewer
                className="min-h-[250px]"
                key={`${Math.random()}`}
                fontSize={11}
                script={
                  {
                    script_type: lang,
                    source_code: (children as string) || "",
                  } as CodeSchema
                }
              />
            );
          },
        }}
      >
        {props.markdown}
      </Markdown>
    </div>
  );
};

export default MarkdownViewer;
