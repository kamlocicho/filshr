import { Flex, Heading } from "@adobe/react-spectrum";
import AppLink from "../AppLink";

export default function Navbar() {
    return (
        <Flex alignItems={"center"} height={"100%"} marginX={"size-200"} gap={"size-200"}>
            <Heading level={2} UNSAFE_className="text-white">Filshr</Heading>
            <AppLink mode="dark" to="/">Home</AppLink>
        </Flex>
    )
}
