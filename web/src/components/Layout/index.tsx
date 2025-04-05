import { Grid, View } from '@adobe/react-spectrum'
import Navbar from './Navbar'
import { Outlet, useNavigation } from 'react-router-dom'
import Loading from '../Loading';

export default function Layout() {
    const navigation = useNavigation();
    const isNavigating = Boolean(navigation.location)

    return (
        <Grid
            areas={[
                'header',
                'content',
                'footer'
            ]}
            height={"100%"}
            rows={['size-1000', 'auto', 'size-1000']}>
            <View backgroundColor={"gray-800"} gridArea={"header"}>
                <Navbar />
            </View>
            <View backgroundColor={"gray-200"} gridArea={"content"}>
                {isNavigating
                    ? <Loading />
                    : <Outlet />}
            </View>
            <View backgroundColor={"static-white"} gridArea={"footer"} />
        </Grid>
    )
}
