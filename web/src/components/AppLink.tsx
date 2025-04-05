import { LinkProps, Link as RouterLink } from 'react-router-dom';

interface IAppLink extends LinkProps {
    mode: "light" | "dark"
}

export default function AppLink({ mode = "light", children, ...props }: IAppLink) {
    return (
        <RouterLink className={`link ${mode == "dark" ? "text-white" : ""}`} {...props}>{children}</RouterLink>
    )
}
