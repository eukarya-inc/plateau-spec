# D.6.2.6 TrafficArea、AuxiliaryTrafficAreaを構成するMultiSurfaceの作成単位

## D.6.2.6 TrafficArea、AuxiliaryTrafficAreaを構成するMultiSurfaceの作成単位

TrafficArea、AuxiliaryTrafficAreaを構成するMultiSurfaceは同一機能（車道部、歩道部など）となるSurfaceの集まりを1つのMultiSurfaceとして作成する。

ただし、機能が車線となる場合、隣接する別の車線とも同一機能が連続するが、車線ごとに区別する必要があるため、別々のTrafficAreaとして作成する。

また、同一機能を有するSurfaceが連続していない場合は、別々のTrafficArea、AuxiliaryTrafficAreaとしてよい。

![Image](images/_trafficareaauxiliarytrafficareaを構成する_multisurfaceの作成単位__img001.png)

